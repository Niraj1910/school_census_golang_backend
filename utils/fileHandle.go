package utils

import (
	"context"
	"fmt"
	"mime/multipart" // Add this import
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveFileLocally(c *gin.Context, fileHeader *multipart.FileHeader) error {
	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll("assets/uploads", 0755); err != nil {
		return fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Generate unique filename
	// uniqueFilename := generateUniqueFilename(fileHeader.Filename)
	localFilePath := filepath.Join("assets/uploads", fileHeader.Filename)

	// Save file locally
	if err := c.SaveUploadedFile(fileHeader, localFilePath); err != nil {
		return fmt.Errorf("failed to save file locally: %v", err)
	}

	return nil
}

func UploadToCloudinary(fileHeader *multipart.FileHeader, schoolName, emailAddress, contactNumber string) (string, error) {
	// Generate public ID
	publicID := generatePublicID(schoolName, emailAddress, contactNumber)

	// Remove from local after function completes
	defer func() {
		os.Remove("assets/uploads/" + fileHeader.Filename)
	}()

	// Get Cloudinary configuration from environment variables
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	// Create Cloudinary URL
	cloudinaryURL := fmt.Sprintf("cloudinary://%s:%s@%s", apiKey, apiSecret, cloudName)

	// Initialize Cloudinary
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", fmt.Errorf("failed to initialize Cloudinary: %v", err)
	}

	// Upload to Cloudinary
	ctx := context.Background()
	uploadParams := uploader.UploadParams{
		PublicID:     publicID,
		Folder:       "schoolImages",
		ResourceType: "image",
	}

	resp, err := cld.Upload.Upload(ctx, "assets/uploads/"+fileHeader.Filename, uploadParams)
	if err != nil {
		return "", fmt.Errorf("cloudinary upload failed: %v", err)
	}

	return resp.SecureURL, nil
}

func generateUniqueFilename(originalFilename string) string {
	extension := filepath.Ext(originalFilename)
	timestamp := time.Now().Format("20060102150405")
	uniqueID := uuid.New().String()[:8]

	return fmt.Sprintf("%s_%s%s", timestamp, uniqueID, extension)
}

func generatePublicID(schoolName, emailAddress, contactNumber string) string {
	publicID := fmt.Sprintf("%s_%s_%s", schoolName, emailAddress, contactNumber)

	// Clean the public ID (remove special characters)
	publicID = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, publicID)

	return publicID
}
