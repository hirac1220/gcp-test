// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
)

func main() {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dlBucketName := os.Getenv("DL_BUCKET")
	dlFileName := os.Getenv("DL_FILE_NAME")
	localFileName := os.Getenv("LOCAL_FILE_NAME")

	// Create the local file
	flocal, err := os.Create(localFileName)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer flocal.Close()

	// Download the sample file
	var data []byte
	if data, err = Get(dlBucketName, dlFileName); err != nil {
		log.Fatalf("Failed to download data from bucket: %v", err)
	}
	// Write to the local file
	flocal.Write(([]byte)(data))

	return
}

func Get(bucket, path string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	r, err := client.Bucket(bucket).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
