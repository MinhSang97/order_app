package errors

import "errors"

var (
	//UserConflict   = errors.New("Lỗi: người dùng đã tồn tại")
	CreatOTPFail = errors.New("Gửi OTP thất bại")
	OTPVerified  = errors.New("OTP code không tồn tại")
	//UserNotFound   = errors.New("Email này chưa được đăng kí")
	//UserNotUpdated = errors.New("Cập nhật không thành công")
	//UserNotDeleted = errors.New("Xoá không thành công")
	//NotAdmin       = errors.New("Không phải quyền Admin")
	//NotUsers       = errors.New("Không phải quyền Users")
)
