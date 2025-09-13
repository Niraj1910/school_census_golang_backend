package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/Niraj1910/school-census-go-backend/config"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadSchoolImage(c *gin.Context) (string, error) {

	file, fileHeader, err := c.Request.FormFile("schoolImage")
	if err != nil {
		return "", fmt.Errorf("no image file found: %v", err)
	}
	defer file.Close()

	if !isImageFile(fileHeader) {
		return "", fmt.Errorf("invalid image file type: %s", fileHeader.Header.Get("Content-Type"))
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	schoolName := c.PostForm("schoolName")
	emailAddress := c.PostForm("emailAddress")
	contactNumber := c.PostForm("contactNumber")

	publicID := fmt.Sprintf("%s_%s_%s", schoolName, emailAddress, contactNumber)
	publicID = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, publicID)

	// Upload to Cloudinary
	ctx := context.Background()

	uploadResult, err := config.Cloudinary.Upload.Upload(ctx, fileBytes, uploader.UploadParams{
		Folder:       "schoolImages",
		PublicID:     publicID,
		ResourceType: "image",
	})

	if err != nil {
		return "", fmt.Errorf("cloudinary upload failed: %v", err)
	}

	return uploadResult.SecureURL, nil
}

func isImageFile(fileHeader *multipart.FileHeader) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	return strings.HasPrefix(contentType, "image/")
}
