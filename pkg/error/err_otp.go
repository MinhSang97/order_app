package errors

import "errors"

var (
	CreatOTPFail            = errors.New("Gửi OTP thất bại")
	OTPVerified             = errors.New("OTP code không tồn tại")
	ChangePasswordByOTPFail = errors.New("Thay đổi password bới OTP thất bại")
)
