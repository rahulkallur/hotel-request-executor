package services

import (
	"Executor/constant"
	model "Executor/models"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

var (
	producer      sarama.SyncProducer
	listeners     = make(map[string]chan string)
	listenersLock sync.Mutex
)

type ExecutorRepository interface {
	ExecutorRequest(ExecutorRQ string) string
}

type executorRepository struct{}

func NewExecutorRequestRepository() ExecutorRepository {
	return &executorRepository{}
}

// ExecutorRequest sends a request to the API and returns the response as a string.
func (r *executorRepository) ExecutorRequest(ExecutorRQ string) string {
	var RS = model.Resp{}
	var roomInfo []model.RoomInfo
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(ExecutorRQ), &data); err != nil {
		fmt.Println("Error unmarshaling:", err)
	}

	supplierRequestRaw := data["supplier_request"].(string)

	client := &http.Client{
		Timeout: constant.Timeout,
	}

	RS.RoomInfo = data["room_info"].(string)
	if err := json.Unmarshal([]byte(RS.RoomInfo), &roomInfo); err != nil {
		fmt.Println("Error unmarshaling:", err)
	}

	httpRequestMessage, err := CreateHttpRequestMessage(os.Getenv("API_URL"), supplierRequestRaw)
	if err != nil {
		return err.Error()
	}

	resp, err := client.Do(httpRequestMessage)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return err.Error()
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	if resp.StatusCode != http.StatusOK {
		return err.Error()
	}

	responseData, err := io.ReadAll(reader)
	if err != nil {
		return err.Error()
	}

	hotelresp := string(responseData)
	RS.SupplierResp = strings.Replace(hotelresp, "&", "", -1)

	executor_resp, _ := json.Marshal(RS)
	return string(executor_resp)
}

// CreateHttpRequestMessage creates and configures a new HTTP POST request.
func CreateHttpRequestMessage(url, request string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(request)))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", os.Getenv("MediaType"))
	req.Header.Set("Accept", os.Getenv("MediaType"))
	req.Header.Set("Accept-Encoding", os.Getenv("Gip"))
	req.Header.Set(constant.APIKeyHeaderName, os.Getenv("UserName"))
	req.Header.Set(constant.SignatureHeaderName, CreateSignature())

	return req, nil
}

// CreateSignature generates a signature using the API key, shared secret, and current timestamp.
func CreateSignature() string {
	ts := time.Now().Unix()
	hashString := fmt.Sprintf("%s%s%d", os.Getenv("UserName"), os.Getenv("Password"), ts)

	hash := sha256.New()
	hash.Write([]byte(hashString))

	return hex.EncodeToString(hash.Sum(nil))
}

// Kafka Consumer Group Handler
type ConsumerGroupHandler struct{}

func (h *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("ConsumerGroupHandler Setup")
	return nil
}

func (h *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("ConsumerGroupHandler Cleanup")
	return nil
}

func (h *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Printf("Consuming messages from partition %d\n", claim.Partition())
	for msg := range claim.Messages() {
		processMessages(string(msg.Value))
		session.MarkMessage(msg, "")
	}
	return nil
}

// DataConsumer initializes the Kafka consumer group.
func DataConsumer() {

	topic := os.Getenv("KafkaTopicName")
	brokers := []string{os.Getenv("KafkaBrokerUrl")}

	// Setup a new consumer group
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	consumerGrp := "my-group" + strconv.FormatInt(currentTime, 10)

	// Create a new consumer group
	consumerGroup, err := sarama.NewConsumerGroup(brokers, consumerGrp, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// Handle system signals for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Start consuming messages
	handler := &ConsumerGroupHandler{}
	for {
		if err := consumerGroup.Consume(context.Background(), []string{topic}, handler); err != nil {
			log.Printf("Error while consuming messages: %v", err)
		}

		// Exit if an interrupt signal is received
		select {
		case <-signals:
			log.Println("Interrupt signal received. Shutting down consumer.")
			return
		default:
		}
	}
}

// processMessages is used to handle individual messages received from Kafka
func processMessages(message string) {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	log.Printf("Message received at: %d %d", currentTime)
	fmt.Println("Received message:", message)

	// Struct to hold the parsed response
	var outerData map[string]interface{}
	var innerData map[string]interface{}
	var RS model.Resp
	var roomInfo []model.RoomInfo

	// Unmarshal the outer JSON message
	if err := json.Unmarshal([]byte(message), &outerData); err != nil {
		fmt.Println("Error unmarshaling outer message:", err)
	}

	// Extract correlationId from the outerData
	correlationId, ok := outerData["correlationId"].(string)
	if !ok {
		fmt.Println("Error: 'correlationId' key is missing or not a string")
	}

	// Safely retrieve the "data" key as a string
	dataRaw, ok := outerData["data"].(string)
	if !ok {
		fmt.Println("Error: 'data' key is missing or not a string")
	}

	// Unmarshal the nested "data" JSON string
	if err := json.Unmarshal([]byte(dataRaw), &innerData); err != nil {
		fmt.Println("Error unmarshaling 'data':", err)
		return
	}

	// Safely retrieve and validate "supplier_request"
	supplierRequestRaw, ok := innerData["supplier_request"].(string)
	if !ok {
		fmt.Println("Error: 'supplier_request' key is missing or not a string")
		return
	}

	// Safely retrieve and validate "room_info"
	roomInfoRaw, ok := innerData["room_info"].(string)
	if !ok {
		fmt.Println("Error: 'room_info' key is missing or not a string")
	}

	// Unmarshal "room_info" into the RoomInfo slice
	if err := json.Unmarshal([]byte(roomInfoRaw), &roomInfo); err != nil {
		fmt.Println("Error unmarshaling 'room_info':", err)
	}

	// Assign the parsed roomInfo to RS
	RS.RoomInfo = roomInfoRaw

	// Create HTTP client
	client := &http.Client{
		Timeout: constant.Timeout,
	}

	apiStartTime := time.Now().UnixNano() / int64(time.Millisecond)
	// apiStartMiliseconds := apiStartTime % 1000
	// log.Printf("API request started at: %d %d", apiStartTime, apiStartMiliseconds)

	// Create HTTP request using supplierRequestRaw
	httpRequestMessage, err := CreateHttpRequestMessage(os.Getenv("API_URL"), supplierRequestRaw)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
	}

	// Execute HTTP request
	resp, err := client.Do(httpRequestMessage)
	if err != nil {
		fmt.Println("Error executing HTTP request:", err)

	}
	defer resp.Body.Close()

	// Handle response decompression if needed
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println("Error creating gzip reader:", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// Ensure HTTP response status is OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP request failed with status code %d\n", resp.StatusCode)
	}

	apiEndTime := time.Now().UnixNano() / int64(time.Millisecond)
	// apiEndMiliseconds := apiEndTime % 1000

	log.Printf("API request ended at: %d", apiEndTime-apiStartTime)
	// Read response body
	responseData, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	// Clean up response data and assign to RS
	hotelresp := string(responseData)
	RS.SupplierResp = strings.ReplaceAll(hotelresp, "&", "")

	// Marshal final response and log it
	executorResp, err := json.Marshal(RS)
	if err != nil {
		fmt.Println("Error marshaling executor response:", err)
	}

	// fmt.Println("Executor response:", string(correlationId))

	// Process response to readable format
	repo := NewSearchResponseRepository()
	hotelResponse := repo.SearchResponsetMapper(string(executorResp))

	// fmt.Println("Hotel response:", hotelResponse)

	hotelResponseJSON, err := json.Marshal(hotelResponse)
	if err != nil {
		fmt.Println("Error marshaling hotel response:", err)
		return
	}

	// fmt.Println("Hotel response JSON:", string(hotelResponseJSON))

	// Log the final JSON response
	// fmt.Println("Hotel response JSON:", string(hotelResponseJSON))

	// Publish response to Kafka topic
	err = publishMessage(os.Getenv("PublishTopicName"), string(correlationId), hotelResponseJSON)
	if err != nil {
		fmt.Println("Error publishing message:", err)
	}
}

// init initializes the Kafka producer.
func init() {
	producer = initProducer()
}

func initProducer() sarama.SyncProducer {
	brokerURL := constant.KafkaBrokerUrl
	if brokerURL == "" {
		log.Fatal("KAFKA_BROKER_URL environment variable is not set")
	}

	brokers := strings.Split(brokerURL, ",")

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	log.Printf("Kafka producer initialized successfully with brokers: %v", brokers)
	return producer
}

func publishMessage(topic, key string, message []byte) error {
	if producer == nil {
		log.Println("Producer is not initialized")
		return fmt.Errorf("producer is nil")
	}

	if topic == "" {
		log.Println("Topic is empty")
		return fmt.Errorf("topic is empty")
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(message),
	}
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to topic %s: %v", topic, err)
	}
	return err
}
