
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
		log.Printf("Created bucket %s", "parser")
	}

	// Routes
	router := httprouter.New()
	// Return minio presigned URLs
	router.GET("/models", Cors(GetModels))
	router.GET("/model", Cors(GetModel))
	router.PUT("/model", Cors(UploadModel))
	router.GET("/data", Cors(GetData))
	router.PUT("/data", Cors(UploadData))
	router.GET("/labels", Cors(GetLabels))
	router.PUT("/labels", Cors(UploadLabels))
	router.GET("/data/batch", Cors(GetBatchData))
	router.GET("/labels/batch", Cors(GetBatchLabels))
	router.GET("/data_parser", Cors(GetDataParser))
	router.PUT("/data_parser", Cors(UploadDataParser))
  router.GET("/batch", Cors(GetBatch))
	router.POST("/batch", Cors(BatchData))
	router.GET("/metadata", Cors(GetMetadata))
	router.PUT("/metadata", Cors(UploadMetadata))

	router.POST("/parse", TestParse)

	// Start server
	log.Printf("starting server on %s", listen)
	log.Fatal(http.ListenAndServe(listen, router))
}

func Cors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, p)
	}
}

type BucketsInfo struct {
	Models []string `json:"models"`
}

func GetModels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	bucketNames := BucketsInfo{}
	bucketNames.Models = make([]string, 0)
	for _, bucket := range buckets {
		bucketNames.Models = append(bucketNames.Models, bucket.Name)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bucketNames)
}

func UploadModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bucketName := r.FormValue("id")

	// Create bucket if it doesn't exist
	err := minioClient.MakeBucket(bucketName, "us-east-1")
	if err != nil {
		exists, err := minioClient.BucketExists(bucketName)
		if !(err == nil && exists) {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	expiry := time.Second * 120
	presignedURL, err := minioClient.PresignedPutObject(bucketName, "model", expiry)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
