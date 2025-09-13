package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Niraj1910/school-census-go-backend/config"
	"github.com/Niraj1910/school-census-go-backend/model"
	"github.com/Niraj1910/school-census-go-backend/utils"
	"github.com/gin-gonic/gin"
)

func GetSchools(c *gin.Context) {

	var schools []model.School

	if err := config.DB.Find(&schools).Error; err != nil {
		log.Printf("Database error in GetSchools: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve schools"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": schools})
}

func GetSchoolByID(c *gin.Context) {

	id := c.Param("id")
	var school model.School

	result := config.DB.First(&school, id)
	if result.Error != nil {
		log.Printf("Database error in GetSchoolByID: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No records found with the id"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No school found with the passed id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": school})
}

func CreateSchool(c *gin.Context) {

	if err := c.Request.ParseMultipartForm(20 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse form: " + err.Error()})
		return
	}

	fmt.Println("Form parsed successfully")
	fmt.Println("Form values:", c.Request.PostForm)

	// Check if file exists
	file, fileHeader, err := c.Request.FormFile("schoolImage")
	if err != nil {
		fmt.Println("No schoolImage file found:", err)
	} else {
		defer file.Close()
		fmt.Printf("File received: %s, Size: %d bytes, Type: %s\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header.Get("Content-Type"))
	}

	// Get form values
	schoolName := c.PostForm("schoolName")
	emailAddress := c.PostForm("emailAddress")
	address := c.PostForm("address")
	city := c.PostForm("city")
	state := c.PostForm("state")
	contactNumber := c.PostForm("contactNumber")

	fmt.Printf("Form data: %s, %s, %s, %s, %s, %s\n",
		schoolName, emailAddress, address, city, state, contactNumber)

	var imagePath string
	var uploadErr error

	// Check if image was uploaded
	if _, _, err := c.Request.FormFile("schoolImage"); err == nil {
		imagePath, uploadErr = utils.UploadSchoolImage(c)
		if uploadErr != nil {
			fmt.Println("Image upload error:", uploadErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image: " + uploadErr.Error()})
			return
		}
		fmt.Println("Image uploaded successfully:", imagePath)
	}

	// Create school model
	school := model.School{
		Name:    schoolName,
		Email:   emailAddress,
		Address: address,
		City:    city,
		State:   state,
		Contact: contactNumber,
		Image:   imagePath, // This can be empty if no image was uploaded
	}

	if err := config.DB.Create(&school).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create school: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "School created successfully",
		"data":    school,
	})
}

func UpdateSchool(c *gin.Context) {
	id := c.Param("id")
	var school model.School
	if err := config.DB.First(&school, id).Error; err != nil {
		log.Printf("Database error in UpdateSchool (find): %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	// Parse multipart form for file upload
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse form"})
		return
	}

	// Get form values
	schoolName := c.PostForm("schoolName")
	emailAddress := c.PostForm("emailAddress")
	address := c.PostForm("address")
	city := c.PostForm("city")
	state := c.PostForm("state")
	contactNumber := c.PostForm("contactNumber")

	var imagePath string
	var uploadErr error

	if _, _, err := c.Request.FormFile("schoolImage"); err == nil {
		imagePath, uploadErr = utils.UploadSchoolImage(c)
		if uploadErr != nil {
			fmt.Println("Image upload error:", uploadErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image: " + uploadErr.Error()})
			return
		}
		fmt.Println("Image uploaded successfully:", imagePath)
	}

	// Update fields
	school.Name = schoolName
	school.Email = emailAddress
	school.Address = address
	school.City = city
	school.State = state
	school.Contact = contactNumber
	if imagePath != "" {
		school.Image = imagePath
	}

	if err := config.DB.Save(&school).Error; err != nil {
		log.Printf("Database error in UpdateSchool (save): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update school"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": school})
}

func DeleteSchool(c *gin.Context) {
	id := c.Param("id")
	var school model.School
	if err := config.DB.First(&school, id).Error; err != nil {
		log.Printf("Database error in DeleteSchool (find): %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}
	if err := config.DB.Delete(&school).Error; err != nil {
		log.Printf("Database error in DeleteSchool (delete): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete school"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "School deleted successfully"})
}
