package repo

import (
	"context"
	"github.com/MinhSang97/order_app/model/admin_model"
)

type AdminFunctionRepo interface {
	//GetOneByID(ctx context.Context, id int) (admin_model.AdminFunction, error)
	GetAll(ctx context.Context) ([]admin_model.AdminFunctionModel, error)
	//UpdateOne(ctx context.Context, id int, student *model.Student) error
	//DeleteOne(ctx context.Context, id int) error
	//Search(ctx context.Context, value string) ([]model.Student, error)
	//CreateStudent(ctx context.Context, student *model.Student) error
}
