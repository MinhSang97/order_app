package repo

import (
	"context"
	"github.com/MinhSang97/order_app/model"
)

type OtpRepo interface {
	SendOtp(ctx context.Context, otp *model.OtpModel) error
	//GetAdmin(ctx context.Context, admin *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error)
	//UpdateAdmin(ctx context.Context, user_id string, admin *admin_model.Admin) error
	//DeleteAdmin(ctx context.Context, user_id string) error
}