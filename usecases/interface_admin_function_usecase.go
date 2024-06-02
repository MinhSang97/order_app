package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/model/admin_model"
)

type AdminFunctionUsecase interface {
	GetAll(ctx context.Context) ([]admin_model.AdminFunctionModel, error)
	Edit(ctx context.Context, user_id string, users *admin_model.AdminFunctionModel) error
	AddUser(ctx context.Context, users *admin_model.AdminFunctionModel) error
	DeleteUsers(ctx context.Context, email string) error
}
