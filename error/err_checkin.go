package errors

import "errors"

var (
	CheckInFail  = errors.New("Lỗi: Check-in thất bại")
	CheckOutFail = errors.New("Lỗi: Check-out thất bại")
	HistoryFail  = errors.New("Lỗi: Lấy lịch sử thất bại")
	//SignUpFail     = errors.New("Đăng kí thất bại")
	//UserNotFound   = errors.New("Email này chưa được đăng kí")
	//UserNotUpdated = errors.New("Cập nhật không thành công")
	//UserNotDeleted = errors.New("Xoá không thành công")
)
