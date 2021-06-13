package main

import (
	"log"
	"os"

	"get.cutie.cafe/rainy/server"
	"get.cutie.cafe/rainy/upload"
	"get.cutie.cafe/rainy/upload/filesystem"
	_ "github.com/joho/godotenv/autoload"
)

var uploader upload.Uploader

func main() {
	var err error

	switch os.Getenv("RAINY_UPLOADER") {
	case "filesystem":
		log.Printf("Creating filesystem uploader at %s", os.Getenv("RAINY_UPLOADER_PATH"))
		uploader, err = filesystem.New(os.Getenv("RAINY_UPLOADER_PATH"))
	default:
		log.Fatalf("An uploader system must be defined. See RAINY_UPLOADER")
	}

	if err != nil {
		panic(err)
	}

	server.New(&uploader).Listen("0.0.0.0:4000")
}
