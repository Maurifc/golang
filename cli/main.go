package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

const DEBUG = true

var projectId string
var bucketName string
var keyPath string

func main() {
	if len(os.Args) < 4 {
		fmt.Println("TODO: Print usage")
		os.Exit(0)
	}

	projectId = os.Args[1]
	bucketName = os.Args[2]
	keyPath = os.Args[3]

	if DEBUG {
		fmt.Println("Project ID:", projectId)
		fmt.Println("Bucket name:", bucketName)
		fmt.Println("Key path:", keyPath)
	}

	downloadBucketFiles()
}

func downloadBucketFiles() error {
	// Context
	ctx := context.Background()

	// Get Client
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Bucket instance
	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	// Iterate over all object in the Bucket
	it := bucket.Objects(ctx, nil)

	for {
		attrs, err := it.Next() // Get object information (attrs)

		if err == iterator.Done {
			break
		}

		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucketName, err)
		}

		// Get Object Reader
		rc, err := bucket.Object(attrs.Name).NewReader(ctx)

		if err != nil {
			return fmt.Errorf("Object(%q).NewReader: %v", attrs.Name, err)
		}

		defer rc.Close()

		// Create destination file (empty)
		file, err := os.Create(attrs.Name)

		if err != nil {
			fmt.Errorf("Failed when creating file:", err)
		}
		defer file.Close()

		// Copy bytes from Bucket to file
		_, err = io.Copy(file, rc)

		if err != nil {
			fmt.Errorf("io.Copy: %v", err)
		}

		file.Close()

		fmt.Printf("Object %v downloaded to local file %v\n", attrs.Name, file.Name())
	}

	return nil
}
