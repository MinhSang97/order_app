package usecases

import (
	"context"
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/model/users_model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/MinhSang97/order_app/repo/postgres"
)

type usersUseCase struct {
	usersRepo repo.UsersRepo
}

func NewUsersUseCase() UsersUsecase {
	db := dbutil.ConnectDB()
	usersRepo := postgres.NewUsersRepository(db)
	return &usersUseCase{
		usersRepo: usersRepo,
	}
}

func (uc *usersUseCase) CreateUsers(ctx context.Context, users *users_model.Users) error {
	return uc.usersRepo.CreateUsers(ctx, users)
}
func (uc *usersUseCase) GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error) {
	return uc.usersRepo.GetUsers(ctx, users)
}

func (uc *usersUseCase) UpdateUsers(ctx context.Context, user_id string, users *users_model.Users) error {
	return uc.usersRepo.UpdateUsers(ctx, user_id, users)
}

func (uc *usersUseCase) DeleteUsers(ctx context.Context, user_id string) error {
	return uc.usersRepo.DeleteUsers(ctx, user_id)
}

// UserFunction
func (uc *usersUseCase) GetAddressUsersFunction(ctx context.Context, user_id string) (*users_model.UsersAddressModel, error) {
	return uc.usersRepo.GetAddressUsersFunction(ctx, user_id)
}

func (uc *usersUseCase) AddAddressUsersFunction(ctx context.Context, user_id string, address *users_model.UsersAddressModel) error {
	return uc.usersRepo.AddAddressUsersFunction(ctx, user_id, address)
}

func (uc *usersUseCase) DefaultAddressUsersFunction(ctx context.Context, user_id string, address *users_model.UsersAddressModel) error {
	return uc.usersRepo.DefaultAddressUsersFunction(ctx, user_id, address)
}

// UsersOrder
func (uc *usersUseCase) AddOrderUsersOrder(ctx context.Context, user_id string, order *model.OrderModel) (*model.OrderModel, error) {
	return uc.usersRepo.AddOrderUsersOrder(ctx, user_id, order)
}

func (uc *usersUseCase) StatusOrderUserOrder(ctx context.Context, user_id string, order *model.OrderModel) error {
	return uc.usersRepo.StatusOrderUserOrder(ctx, user_id, order)
}

func (uc *usersUseCase) HistoryOrderUserOrder(ctx context.Context, user_id string) ([]users_model.ResOrderHistory, error) {
	return uc.usersRepo.HistoryOrderUserOrder(ctx, user_id)
}
