package usecases

import (
	"github.com/MinhSang97/order_app/dbutil"
	"github.com/MinhSang97/order_app/model/users_model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/MinhSang97/order_app/repo/mysql"
	"context"
)

type usersUseCase struct {
	usersRepo repo.UsersRepo
}

func NewUsersUseCase() UsersUsecase {
	db := dbutil.ConnectDB()
	usersRepo := mysql.NewUsersRepository(db)
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
