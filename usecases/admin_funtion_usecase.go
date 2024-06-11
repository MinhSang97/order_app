package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/model/admin_model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/MinhSang97/order_app/repo/postgres"
)

type adminFunctionUseCase struct {
	adminFunctionRepo repo.AdminFunctionRepo
}

func NewAdminFunctionUseCase() AdminFunctionUsecase {
	db := dbutil.ConnectDB()
	adminFunctionRepo := postgres.NewAdminFunctionUseCase(db)
	return &adminFunctionUseCase{
		adminFunctionRepo: adminFunctionRepo,
	}
}

func (uc *adminFunctionUseCase) GetAll(ctx context.Context) ([]admin_model.AdminFunctionModel, error) {
	return uc.adminFunctionRepo.GetAll(ctx)
}

func (uc *adminFunctionUseCase) Edit(ctx context.Context, user_id string, users *admin_model.AdminFunctionModel) error {
	return uc.adminFunctionRepo.Edit(ctx, user_id, users)
}

func (uc *adminFunctionUseCase) AddUser(ctx context.Context, user *admin_model.AdminFunctionModel) error {
	return uc.adminFunctionRepo.AddUser(ctx, user)
}

func (uc *adminFunctionUseCase) DeleteUsers(ctx context.Context, email string) error {
	return uc.adminFunctionRepo.DeleteUsers(ctx, email)
}

// admin_function_menu
func (uc *adminFunctionUseCase) GetMenuAll(ctx context.Context) ([]model.MenuItemsModel, error) {
	return uc.adminFunctionRepo.GetMenuAll(ctx)
}

func (uc *adminFunctionUseCase) AddMenu(ctx context.Context, menu *model.MenuItemsModel) (*model.MenuItemsModel, error) {
	return uc.adminFunctionRepo.AddMenu(ctx, menu)
}

func (uc *adminFunctionUseCase) EditMenu(ctx context.Context, item_id string, menu *model.MenuItemsModel) error {
	return uc.adminFunctionRepo.EditMenu(ctx, item_id, menu)
}

func (uc *adminFunctionUseCase) DeleteMenu(ctx context.Context, item_id string) error {
	return uc.adminFunctionRepo.DeleteMenu(ctx, item_id)
}
