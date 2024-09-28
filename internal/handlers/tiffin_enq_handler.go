// tiffin_enq_handler.go
package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"tg/internal/models"
	"tg/internal/services"
)

type TiffinEnqHandler struct {
	service service.TiffinEnqService
}

func NewTiffinEnqHandler(service service.TiffinEnqService) *TiffinEnqHandler {
	return &TiffinEnqHandler{
		service: service, // Assign the service
	}
}

func (h *TiffinEnqHandler) CreateEnquiry(c *gin.Context) {
	var err error

	name := c.Param("name")
	mobile := c.Param("mobile")
	location := c.Param("location")
	startDateStr := c.Param("start_date")

	var input models.TiffinEnq

	// bind input JSON
	// Bind the request body to the input model
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("input %+v\n", input)

	// Validate empty fields
	if name == "" || mobile == "" || location == "" || startDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (name, mobile, location, starting date) are required."})
		return
	}
	// Validate mobile (must be 10 digits)
	mobileRegex := regexp.MustCompile(`^[0-9]{10}$`)
	if !mobileRegex.MatchString(mobile) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number must be 10 digits."})
		return
	}
	// Parse the start_date and validate that it's a valid date and greater than the current date
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid starting date format. Use YYYY-MM-DD."})
		return
	}
	// Check if the start date is greater than the current date
	if startDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "starting date must be greater than the current date."})
		return
	}

	// If all validations pass
	c.JSON(http.StatusOK, gin.H{
		"name":       name,
		"mobile":     mobile,
		"location":   location,
		"start_date": startDate,
		"message":    "Enquiry submitted successfully.",
	})
}

// GetEnquiry retrieves a specific enquiry by ID
func (h *TiffinEnqHandler) GetEnquiry(c *gin.Context) {
	id := c.Param("id")

	// Convert id to integer
	enquiryID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enquiry ID"})
		return
	}

	// Retrieve enquiry from service
	enquiry, err := h.service.GetTiffinEnq(enquiryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enquiry not found"})
		return
	}

	// Return the enquiry details
	c.JSON(http.StatusOK, gin.H{"enquiry": enquiry})
}
