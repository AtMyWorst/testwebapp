
package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
  "math/big"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/minio/minio-go"
	"github.com/yuin/gopher-lua"
)

var listen string
var minioClient *minio.Client

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	listen = os.Getenv("LISTEN")
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	minioID := os.Getenv("MINIO_ACCESS_KEY")
	minioKey := os.Getenv("MINIO_SECRET_KEY")

	// Minio client
	minioClient, err = minio.New(minioEndpoint, minioID, minioKey, false)
	if err != nil {
		log.Fatal("Error loading minio")
	}

	// Create bucket if it doesn't exist
	err = minioClient.MakeBucket("parser", "us-east-1")
	if err != nil {
		exists, err := minioClient.BucketExists("parser")
		if err == nil && exists {
			log.Printf("Bucket %s already exists", "parser")
		} else {
			log.Printf("%s", err)
			log.Fatal("Error creating bucket")
		}
	} else {