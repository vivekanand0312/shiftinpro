package utility

import (
	"sync"
	"time"
)

const OTP_MAX_MINS float64 = 2
const OTP_MAX_ATTEMPT_COUNT int = 3

type OTPInfo struct {
	Attempts int
	LastSent time.Time
}

var otpCache = sync.Map{}

func GetOTPInfo(phone string) *OTPInfo {
	if val, ok := otpCache.Load(phone); ok {
		return val.(*OTPInfo)
	}
	return nil // Return nil if no record found
}

func SetOTPInfo(phone string, info *OTPInfo) {
	otpCache.Store(phone, info)
}

func ResetOTPInfo(phone string) {
	otpCache.Delete(phone)
}
