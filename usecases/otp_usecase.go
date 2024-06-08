package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/MinhSang97/order_app/repo/postgres"
)

type otpUseCase struct {
	otpRepo repo.OtpRepo
}

func NewOtpUseCase() OtpUsecase {
	db := dbutil.ConnectDB()
	otpRepo := postgres.NewOtpRepository(db)
	return &otpUseCase{
		otpRepo: otpRepo,
	}
}

func (uc *otpUseCase) SendOtp(ctx context.Context, otp *model.OtpModel) error {
	return uc.otpRepo.SendOtp(ctx, otp)
}

func (uc *otpUseCase) VerifyOtp(ctx context.Context, otp *model.OtpModel) error {
	return uc.otpRepo.VerifyOtp(ctx, otp)
}

func (uc *otpUseCase) ChangePassword(ctx context.Context, otp_code string, otp *model.OtpModel) error {
	return uc.otpRepo.ChangePassword(ctx, otp_code, otp)
}
