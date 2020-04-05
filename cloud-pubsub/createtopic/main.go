// Sample pubsub-quickstart creates a Google Cloud Pub/Sub topic.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

func main() {
	// Read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("PROJECT_ID")
	// Sets the id for the new topic.
	topicID := "test"

	// Creates a client.
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}
