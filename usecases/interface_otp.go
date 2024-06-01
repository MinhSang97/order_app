package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/model"
)

type OtpUsecase interface {
	SendOtp(ctx context.Context, otp *model.OtpModel) error
	VerifyOtp(ctx context.Context, otp *model.OtpModel) error
	ChangePassword(ctx context.Context, otp_code string, otp *model.OtpModel) error
}
