package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/model"
)

type OtpUsecase interface {
	SendOtp(ctx context.Context, otp *model.OtpModel) error
	VerifyOtp(ctx context.Context, otp *model.OtpModel) error
	//Edit(ctx context.Context, user_id string, users *admin_model.AdminFunctionModel) error
	//UpdateOne(ctx context.Context, id int, student *model.Student) error
	//DeleteOne(ctx context.Context, id int) error
	//Search(ctx context.Context, Value string) ([]model.Student, error)
	//CreateStudent(ctx context.Context, student *model.Student) error
}
