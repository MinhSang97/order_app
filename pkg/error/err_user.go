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
	UserAddressNotFound = errors.New("User không có tồn tại")
	AddAddressFail      = errors.New("Thêm địa chỉ thất bại")
	DefaultAddressFail  = errors.New("Đặt địa chỉ mặc định thất bại")
	AddMenuItemsFail    = errors.New("Thêm menu items thất bại")
	EditMenuItemsFail   = errors.New("Sửa menu items thất bại")
	ItemIDNotFound      = errors.New("Item không tồn tại")
	DeleteMenuItemsFail = errors.New("Xoá menu items thất bại")
)
