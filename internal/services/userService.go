package services

import (
    "errors"
    "time"

    "shiftinpro/internal/models"
    "shiftinpro/internal/repository"
    "shiftinpro/utility"
)

type UserService interface {
    SaveUser(user *models.User) error
    GetUser(phone string) (*models.User, error)
    ValidateOTP(phone string, otp int) (bool, error)
    SendOTP(phone string) (bool, error)
    UpdateAddress(userID int, input models.ReqUpdateAddress) error
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) SaveUser(user *models.User) error {
    return s.repo.CreateUser(user)
}

func (s *userService) GetUser(phone string) (*models.User, error) {
    return s.repo.GetUserByPhone(phone)
}

func (s *userService) ValidateOTP(phone string, otp int) (bool, error) {
    otpInfo := utility.GetOTPInfo(phone)

    if otpInfo != nil {
        if time.Since(otpInfo.LastSent).Minutes() < utility.OTP_MAX_MINS && otpInfo.Attempts >= utility.OTP_MAX_ATTEMPT_COUNT {
            return false, errors.New("too many attempts, try again after 5 minutes")
        }

        if time.Since(otpInfo.LastSent).Minutes() > 5 {
            utility.ResetOTPInfo(phone)
            otpInfo = &utility.OTPInfo{}
        }
    } else {
        otpInfo = &utility.OTPInfo{}
    }

    if otp != 1234 { // Assuming OTP is constant here, replace with actual logic
        otpInfo.Attempts++
        otpInfo.LastSent = time.Now()
        utility.SetOTPInfo(phone, otpInfo)
        return false, errors.New("invalid OTP")
    }

    otpInfo.Attempts = 0
    utility.SetOTPInfo(phone, otpInfo)
    return true, nil
}

func (s userService) SendOTP(phone string) (bool, error) {
    if phone != "" {
        return true, nil
    } else {
        return false, errors.New("OTP Sending failed!")
    }
}

func (s *userService) UpdateAddress(userID int, input models.ReqUpdateAddress) error {
    user := models.User{
        House:       input.House,
        Area:        input.Area,
        Landmark:    input.Landmark,
        SdAddressID: input.SdAddressID,
    }
    return s.repo.UpdateUserAddress(userID, user)
}
