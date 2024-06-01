package repo

import (
	"context"
	"github.com/MinhSang97/order_app/model"
)

type OtpRepo interface {
	SendOtp(ctx context.Context, otp *model.OtpModel) error
	//GetAdmin(ctx context.Context, admin *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error)
	VerifyOtp(ctx context.Context, otp *model.OtpModel) error
	ChangePassword(ctx context.Context, otp_code string, otp *model.OtpModel) error
	//DeleteAdmin(ctx context.Context, user_id string) error
}
