package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "net/http"
    "regexp"
    "strconv"
    "tg/internal/services"
)

const OTP int = 1234

// Regex pattern for validating phone numbers
var phoneRegex = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`) // E.164 format, example: +1234567890

// Regex pattern for validating names (only letters and spaces, 1-50 characters)
var nameRegex = regexp.MustCompile(`^[a-zA-Z\s]{1,50}$`)

// Regex pattern for validating OTP (4 to 6 digits)
var otpRegex = regexp.MustCompile(`^\d{4,6}$`)

// Validator instance
var validate = validator.New()

// Define your request structure
type ReqUserRegister struct {
    Phone string `json:"phone" validate:"required,e164"` // Assuming custom validation for phone
    Name  string `json:"name" validate:"required"`
    OTP   int    `json:"otp" validate:"required,min=1000,max=9999"` // Using int for OTP
}

type UserHandler struct {
    userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }

    user, err := h.userService.GetUser(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

// func (h *UserHandler) Login(c *gin.Context) {

//     idStr := c.Param("phone")
//     id, err := strconv.Atoi(idStr)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone"})
//         return
//     }

//     if inputOtp == OTP {
//         c.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
//     }
// }

// Custom handler for registration
func (h *UserHandler) Register(c *gin.Context) {
    var input ReqUserRegister

    // Bind and validate the JSON body
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input format", "error": err.Error()})
        return
    }

    // Validate using go-playground validator
    if err := validate.Struct(input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Validation failed", "error": err.Error()})
        return
    }
    fmt.Printf("INPUT: %+v\n", input)

    // If validation passes
    c.JSON(http.StatusOK, gin.H{"status": true, "message": "SUCCESS"})
}
