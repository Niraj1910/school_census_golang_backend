package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var Cloudinary *cloudinary.Cloudinary

func ConnectCloudinary() {
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	// Add debug logging
	fmt.Printf("Cloudinary Config - Name: %s, Key: %s\n", cloudName, apiKey)

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		log.Fatal("Cloudinary environment variables are not set properly")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Fatal("Failed to initialize Cloudinary: ", err)
	}

	ctx := context.Background()
	_, err = cld.Admin.Ping(ctx)
	if err != nil {
		log.Fatal("Cloudinary connection test failed: ", err)
	}

	fmt.Println("Cloudinary connected successfully")
	Cloudinary = cld
}
