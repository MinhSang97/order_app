package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/MinhSang97/order_app/repo/mysql"
)

type otpUseCase struct {
	otpRepo repo.OtpRepo
}

func NewOtpUseCase() OtpUsecase {
	db := dbutil.ConnectDB()
	otpRepo := mysql.NewOtpRepository(db)
	return &otpUseCase{
		otpRepo: otpRepo,
	}
}

func (uc *otpUseCase) SendOtp(ctx context.Context, otp *model.OtpModel) error {
	return uc.otpRepo.SendOtp(ctx, otp)
}
