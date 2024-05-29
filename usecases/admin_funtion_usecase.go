package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/model/admin_model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/MinhSang97/order_app/repo/mysql"
)

type adminFunctionUseCase struct {
	adminFunctionRepo repo.AdminFunctionRepo
}

func NewAdminFunctionUseCase() AdminFunctionUsecase {
	db := dbutil.ConnectDB()
	adminFunctionRepo := mysql.NewAdminFunctionUseCase(db)
	return &adminFunctionUseCase{
		adminFunctionRepo: adminFunctionRepo,
	}
}

func (uc *adminFunctionUseCase) GetAll(ctx context.Context) ([]admin_model.AdminFunctionModel, error) {
	return uc.adminFunctionRepo.GetAll(ctx)
}
