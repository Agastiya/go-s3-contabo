package main

import (
	"log"
	"os"

	"github.com/agastiya/go-s3storage/storage"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

var (
	s3Client  *s3.Client
	s3Storage *storage.S3Storage
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Client = InitStorage()
	s3Storage = storage.NewS3Storage(s3Client, os.Getenv("BUCKET_NAME"), os.Getenv("URL_PREVIEW"))
}

func main() {

	buckets, _ := s3Storage.ListBuckets()
	log.Print(buckets)

	bucketFilePath := "sample-file/images.jpg"
	s3Storage.UploadFile(bucketFilePath)

	fileName := "images.jpg" // or full file path

	file, err := s3Storage.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return
	}
	log.Printf("File content: %s", *file)

	s3Storage.DeleteFile(fileName)
}
