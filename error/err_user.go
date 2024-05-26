package errors

import "errors"

var (
	UserConflict   = errors.New("Lỗi: người dùng đã tồn tại")
	SignUpFail     = errors.New("Đăng kí thất bại")
	UserNotFound   = errors.New("Email này chưa được đăng kí")
	UserNotUpdated = errors.New("Cập nhật không thành công")
	UserNotDeleted = errors.New("Xoá không thành công")
)
