# Go S3 Contabo 🚀

This project demonstrates how to interact with Contabo S3 Object Storage using Go. It provides functionality for basic S3 operations like uploading, downloading, and managing files in Contabo S3 buckets.

## 📋 Prerequisites

- Go 1.16 or higher
- Contabo S3 Storage account
- Access Key and Secret Key from Contabo

## 📥 Installation

```bash
git clone https://github.com/agastiya/go-s3-contabo.git
cd go-s3-contabo
go mod download
```

## ⚙️ Configuration

Before running the application, make sure you have your Contabo S3 credentials properly configured:

- S3 Endpoint
- Access Key
- Secret Key
- Region
- Bucket Name

## 🚀 Usage

The project provides functionality for:

- List all buckets in S3
- Read files in S3 buckets
- Uploading files to S3
- Deleting files from S3

## 📁 Project Structure

```
.
├── go.mod
├── go.sum
├── main.go
├── storage.go
└── sample-file/
    ├── file.txt
    └── images.jpg
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.