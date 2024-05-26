package mysql

import (
	errors "github.com/MinhSang97/order_app/error"
	"github.com/MinhSang97/order_app/log"
	"github.com/MinhSang97/order_app/model/admin_model"
	"github.com/MinhSang97/order_app/model/users_model"
	"github.com/MinhSang97/order_app/repo"
	"context"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func (s adminRepository) CreateAdmin(ctx context.Context, admin *admin_model.Admin) error {
	users := admin

	err := s.db.Table("Users").Create(users).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {

			if driverErr.Number == 1062 {
				return errors.UserConflict
			}
		}
		return errors.SignUpFail
	}
	return nil
}

func (s adminRepository) GetAdmin(ctx context.Context, admin *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error) {
	users := admin

	err := s.db.Table("Users").Where("email = ?", users.Email).First(users).Error
	if err != nil {
		//log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		log.Error(err.Error())
		return nil, err
	}
	return users, nil
}

func (s adminRepository) UpdateAdmin(ctx context.Context, user_id string, admin *admin_model.Admin) error {
	users := admin

	err := s.db.Table("Users").Where("user_id = ?", user_id).Updates(users).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {

			if driverErr.Number == 1062 {
				return errors.UserNotUpdated
			}
		}
		return errors.SignUpFail
	}
	return nil
}

func (s adminRepository) DeleteAdmin(ctx context.Context, user_id string) error {
	var user users_model.Users

	// Check if user exists
	if err := s.db.Table("Users").Where("user_id = ?", user_id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.UserNotFound
		}
		return err
	}

	// If user exists, delete the user
	if err := s.db.Table("Users").Where("user_id = ?", user_id).Delete(&user).Error; err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				return errors.UserNotDeleted
			}
		}
		return err
	}

	return nil
}

var instances adminRepository

func NewAdminRepository(db *gorm.DB) repo.AdminRepo {
	if instances.db == nil {
		instances.db = db

	}
	return instances
}
