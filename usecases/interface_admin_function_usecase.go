package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/model/admin_model"
)

type AdminFunctionUsecase interface {
	//admin_function_member
	GetAll(ctx context.Context) ([]admin_model.AdminFunctionModel, error)
	Edit(ctx context.Context, user_id string, users *admin_model.AdminFunctionModel) error
	AddUser(ctx context.Context, users *admin_model.AdminFunctionModel) error
	DeleteUsers(ctx context.Context, email string) error

	//admin_function_menu
	GetMenuAll(ctx context.Context) ([]model.MenuItemsModel, error)
	AddMenu(ctx context.Context, menu *model.MenuItemsModel) (*model.MenuItemsModel, error)
	EditMenu(ctx context.Context, item_id string, menu *model.MenuItemsModel) error
	DeleteMenu(ctx context.Context, item_id string) error
}
