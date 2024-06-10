package errors

import "errors"

var (
	UserConflict        = errors.New("Lỗi: người dùng đã tồn tại")
	SignUpFail          = errors.New("Đăng kí thất bại")
	SignInFail          = errors.New("Đăng nhập thất bại")
	UserNotFound        = errors.New("Email này chưa được đăng kí")
	UserNotUpdated      = errors.New("Cập nhật không thành công")
	UserNotDeleted      = errors.New("Xoá không thành công")
	NotAdmin            = errors.New("Không phải quyền Admin")
	NotUsers            = errors.New("Không phải quyền Users")
	UserAddressNotFound = errors.New("Không tìm thấy địa chỉ người dùng")
	AddAddressFail      = errors.New("Thêm địa chỉ thất bại")
)
