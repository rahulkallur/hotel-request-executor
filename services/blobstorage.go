package services

import (
	"Executor/constant"
	"context"
	"path/filepath"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type BlobHandler interface {
	WriteFile(sid, operation, filename, message string)
}

type blobStorageHandler struct {
	Client    azblob.Client
	Container string
}

func NewBlobStorageHandler() BlobHandler {
	client, _ := azblob.NewClientFromConnectionString(constant.AzureConnectionString, nil)
	return &blobStorageHandler{Client: *client, Container: constant.Container}
}

// File uploads a given string (item) as a byte slice to the specified file path in the container.
func (c *blobStorageHandler) File(filePath string, item string) error {
	// Create a context for the operation (this is required by many Go cloud libraries)
	ctx := context.Background()

	// Convert the item string into a byte slice (the data to be uploaded)
	data := []byte(item)

	// Upload the data to the specified file path in the container
	// c.Client.UploadBuffer is the method that interacts with the blob storage service
	// It takes the context, container name, file path, the data (byte slice), and optional settings (nil in this case)
	_, err := c.Client.UploadBuffer(ctx, c.Container, filePath, data, nil)

	// Return any error encountered during the upload
	return err
}

// WriteFile creates a file path based on the provided parameters and uploads a message to that path.
func (r *blobStorageHandler) WriteFile(sid, operation, filename, message string) {
	// Format the current date as a string in the format: "DD-MM-YYYY"
	datefolder := time.Now().Format("02-01-2006")

	// Build the directory path for the log file, combining the base directory with the date, session ID, and operation
	logDirectory := filepath.Join("HotelBeds_Condense", datefolder, sid, operation)

	// Combine the log directory and filename to create the full file path
	strfilePath := filepath.Join(logDirectory, filename)

	// Call the File method to upload the message (item) to the generated file path
	r.File(strfilePath, message)
}
