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

	bucketFilePath := "sample-file/images.jpg"
	fileName := "images.jpg" // or full file path

	buckets, err := s3Storage.ListBuckets()
	if err != nil {
		log.Fatalf("Error listing buckets: %v", err)
	}
	log.Print(buckets)

	err = s3Storage.UploadFile(bucketFilePath)
	if err != nil {
		log.Fatalf("Error upload file: %v", err)
	}
	log.Print("Suscessfully uploaded file")

	file, err := s3Storage.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return
	}
	log.Printf("File content: %s", *file)

	err = s3Storage.DeleteFile(fileName)
	if err != nil {
		log.Fatalf("Error deleting file: %v", err)
		return
	}
	log.Printf("File %s deleted successfully", fileName)
}
