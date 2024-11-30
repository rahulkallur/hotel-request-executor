package constant

import "time"

const (
	Timeout               = 10 * time.Second
	APIKeyHeaderName      = "Api-key"
	SignatureHeaderName   = "X-Signature"
	MediaType             = "application/json"
	Gip                   = "gzip"
	SearchMethod          = "Search"
	UserName              = "88e8ad8b590362cbe0134c47b6ad351a"
	Password              = "2cb23e90ac"
	Url                   = "https://api.test.hotelbeds.com/hotel-api/1.0/hotels"
	Container             = "CANTAINER-NAME"
	AzureConnectionString = "USE-YOUR-CONNECTIONSTRING"
	KafkaBrokerUrl        = "localhost:9092"
	KafkaTopicName        = "hotelsearch"
	PublishTopicName      = "response-topic"
	KafkaGroupName        = "request-executor"
)
