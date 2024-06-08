package repo

import (
	"context"
	"github.com/MinhSang97/order_app/model/users_model"
)

type UsersRepo interface {
	CreateUsers(ctx context.Context, users *users_model.Users) error
	GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error)
	UpdateUsers(ctx context.Context, user_id string, users *users_model.Users) error
	DeleteUsers(ctx context.Context, user_id string) error

	//UserFunction
	GetAddressUsersFunction(ctx context.Context, user_id string) (*users_model.UsersAddressModel, error)
}
