package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Client := InitStorage()

	ListBuckets(s3Client)

	// bucketFilePath := "sample-file/images.jpg"
	// UploadFile(s3Client, bucketFilePath)

	// fileName := "images.jpg" //or full file path
	// ReadFile(s3Client, fileName)
	// DeleteFile(s3Client, fileName)

}
