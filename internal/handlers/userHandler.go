package handlers

import (
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    // "mime/multipart"
    "net/http"
    "shiftinpro/internal/models"
    "shiftinpro/internal/services"
    "shiftinpro/utility"
)

var validate = validator.New()

type ReqUserRegister struct {
    Phone    string `json:"phone" validate:"required,e164"`
    Name     string `json:"name" validate:"required"`
    OTP      int    `json:"otp" validate:"required,min=1000,max=9999"`
    UserType int    `json:"userType" validate:"required,min=1,max=6"`
    // Image *multipart.FileHeader `form:"image" binding:"omitempty"`
}

type ReqUserLogin struct {
    Phone string `json:"phone" validate:"required,e164"`
    OTP   int    `json:"otp" validate:"required,min=1000,max=9999"`
}

type ReqResendOTP struct {
    Phone string `json:"phone" validate:"required,e164"`
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
        Phone:         input.Phone,
        Name:          input.Name,
        UserType:      input.UserType,
        IsOtpVerified: true,
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
    user, err := h.userService.GetUserByPhone(input.Phone)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"status": false, "next_screen": "register", "message": "User doesn't exists!", "error": err.Error()})
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

func (h *UserHandler) SendOTP(c *gin.Context) {
    var input ReqResendOTP

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input format", "error": err.Error()})
        return
    }

    if err := validate.Struct(input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Validation failed", "error": err.Error()})
        return
    }

    // Fetch user
    user, err := h.userService.GetUserByPhone(input.Phone)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"status": false, "next_screen": "register", "message": "User doesn't exists!", "error": err.Error()})
        return
    }

    otpInfo := utility.GetOTPInfo(user.Phone)
    if otpInfo != nil {
        if time.Since(otpInfo.LastSent).Minutes() < utility.OTP_MAX_MINS && otpInfo.Attempts >= utility.OTP_MAX_ATTEMPT_COUNT {
            c.JSON(http.StatusTooManyRequests, gin.H{"status": false, "message": "Too many attempts, try again after 5 minutes"})
            return
        }

        if time.Since(otpInfo.LastSent).Minutes() > 5 {
            utility.ResetOTPInfo(user.Phone)
            otpInfo = &utility.OTPInfo{}
        }
    } else {
        otpInfo = &utility.OTPInfo{} // Initialize if nil
    }

    otpInfo.Attempts++
    otpInfo.LastSent = time.Now()
    utility.SetOTPInfo(user.Phone, otpInfo)

    isOtpSent, err := h.userService.SendOTP(user.Phone)
    if err != nil || !isOtpSent {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Something went wrong while sending OTP", "error": err.Error()})
        return
    }

    // Resend OTP logic here
    c.JSON(http.StatusOK, gin.H{"status": true, "message": "OTP sent successfully!"})
}

func (h *UserHandler) UpdateAddress(c *gin.Context) {
    var input models.ReqUpdateAddress

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid input format", "error": err.Error()})
        return
    }

    phone, exists := c.Get("phone")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Unauthorized"})
        return
    }

    user, err := h.userService.GetUserByPhone(phone.(string))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid user ID", "error": err.Error()})
        return
    }

    err = h.userService.UpdateAddress(user.ID, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to update address", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": true, "message": "Address updated successfully"})
}
