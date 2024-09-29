package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    // "mime/multipart"
    "net/http"
    "shiftproin/internal/models"
    "shiftproin/internal/services"
    "shiftproin/utility"
)

const OTP int = 1234

var validate = validator.New()

type ReqUserRegister struct {
    Phone string `json:"phone" validate:"required,e164"`
    Name  string `json:"name" validate:"required"`
    OTP   int    `json:"otp" validate:"required,min=1000,max=9999"`
    // Image *multipart.FileHeader `form:"image" binding:"omitempty"`
}
type ReqUserLogin struct {
    Phone string `json:"phone" validate:"required,e164"`
    OTP   int    `json:"otp" validate:"required,min=1000,max=9999"`
}

type UserHandler struct {
    userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
    var input ReqUserRegister

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input format", "error": err.Error()})
        return
    }

    if err := validate.Struct(input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Validation failed", "error": err.Error()})
        return
    }

    user := models.User{
        Phone: input.Phone,
        Name:  input.Name,
        // Image: imagePath,
        // OTP:   input.OTP,
    }

    // Insert user
    if err := h.userService.SaveUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to save user", "error": err.Error()})
        return
    }

    // Validate OTP
    _, err := h.userService.ValidateOTP(user.Phone, input.OTP)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid phone or OTP", "error": err.Error()})
        return
    }

    token, err := utility.GenerateJWT(input.Phone)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to generate token", "error": err.Error()})
        return
    }

    refreshToken, err := utility.GenerateRefreshToken(input.Phone)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to generate refresh token", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  true,
        "message": "SUCCESS",
        "userInfo": gin.H{
            "name": user.Name,
            // "image":        imagePath,
            "token":        token,
            "refreshToken": refreshToken,
        },
    })
}

func (h *UserHandler) Login(c *gin.Context) {
    var input ReqUserLogin

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input format", "error": err.Error()})
        return
    }

    if err := validate.Struct(input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Validation failed", "error": err.Error()})
        return
    }

    // Fetch user
    user, err := h.userService.GetUser(input.Phone)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid phone or OTP", "error": err.Error()})
        return
    }
    valid, err := h.userService.ValidateOTP(input.Phone, input.OTP)
    if err != nil || !valid {
        c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid phone or OTP", "error": err.Error()})
        return
    }

    token, err := utility.GenerateJWT(user.Phone)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to generate token", "error": err.Error()})
        return
    }

    refreshToken, err := utility.GenerateRefreshToken(user.Phone)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to generate refresh token", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":  true,
        "message": "Login successful",
        "userInfo": gin.H{
            "name":         user.Name,
            "token":        token,
            "refreshToken": refreshToken,
            "image":        user.Image,
        },
    })
}
