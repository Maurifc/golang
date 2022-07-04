package main

import (
	"bufio"
	"context"
	"fmt"
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
	//Writer
	file, err := os.OpenFile("fileslist.txt", os.O_CREATE|os.O_RDWR, 0640)

	if err != nil {
		fmt.Errorf("Failed when opening file to write:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Context
	ctx := context.Background()

	// Client
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Bucket instance
	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	defer cancel()

	it := bucket.Objects(ctx, nil)

	for {
		attrs, err := it.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucketName, err)
		}
		fmt.Fprintln(writer, attrs.Name)
		writer.Flush()
	}

	return nil
}
