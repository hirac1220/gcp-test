// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
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
	upBucketName := os.Getenv("UP_BUCKET")
	upFileName := os.Getenv("UP_FILE_NAME")
	localFileName := os.Getenv("LOCAL_FILE_NAME")

	// Open the local file
	flocal, err := os.Open(localFileName)
	if err != nil {
		log.Fatalf("Failed to open data: %v", err)
	}
	defer flocal.Close()

	// The local os.file to byte
	const BUFSIZE = 1024
	buf := make([]byte, BUFSIZE)
	for {
		n, err := flocal.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
	}
	// Put the local file to GCP storage
	if err := Put(upBucketName, upFileName, buf); err != nil {
		log.Fatalf("Failed to upload data to bucket: %v", err)
	}

	return
}

func Put(bucket, path string, data []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	w := client.Bucket(bucket).Object(path).NewWriter(ctx)
	defer w.Close()

	if n, err := w.Write(data); err != nil {
		return err
	} else if n != len(data) {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return nil
}
