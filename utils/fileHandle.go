package utils

import (
	"io"

	"github.com/Niraj1910/school-census-go-backend/config"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadSchoolImage(c *gin.Context, schoolName, emailAddress, contactNumber string) (string, error) {
	file, fileHeader, err := c.Request.FormFile("schoolImage")
	if err != nil || fileHeader.Size == 0 {
		return "", nil // No image uploaded
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		Folder:       "schoolImages",
		PublicID:     schoolName + "_" + emailAddress + "_" + contactNumber,
		ResourceType: "image",
	}
	uploadRes, err := config.Cloudinary.Upload.Upload(c, fileBytes, uploadParams)
	if err != nil {
		return "", err
	}
	return uploadRes.SecureURL, nil
}
