package mysql

import (
	"context"
	"fmt"
	errors "github.com/MinhSang97/order_app/error"
	"github.com/MinhSang97/order_app/log"
	"github.com/MinhSang97/order_app/model/admin_model"
	"github.com/MinhSang97/order_app/model/users_model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func (s adminRepository) CreateAdmin(ctx context.Context, admin *admin_model.Admin) error {
	// Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return errors.SignUpFail
	}

	// Insert into the users table
	err := tx.Table("users").Create(admin).Error
	if err != nil {
		tx.Rollback()
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				return errors.UserConflict
			}
		}
		return errors.SignUpFail
	}

	// Insert into the user_addresses table
	query := `INSERT INTO order_app.user_addresses (user_id, address) VALUES(?, ?);`
	if err := tx.Exec(query, admin.UserId, admin.Address).Error; err != nil {
		tx.Rollback()
		return errors.SignUpFail
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.SignUpFail
	}

	return nil
}

func (s adminRepository) GetAdmin(ctx context.Context, admin *admin_model.ReqSignIn) (*admin_model.ReqSignIn, error) {
	users := admin

	// Step 1: Query the database to get the user's role based on their email
	var role string
	err := s.db.Table("Users").Select("role").Where("email = ?", users.Email).Scan(&role).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		return nil, err
	}

	// Step 2: Check if the role is "admin"
	if role != "admin" {
		return nil, errors.NotAdmin
	}

	// Step 3: Proceed with the rest of the function if the role is "admin"
	err = s.db.Table("Users").Where("email = ?", users.Email).Updates(users).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 {
				return users, errors.UserNotUpdated
			}
		}
		return users, errors.UserNotUpdated
	}

	err = s.db.Table("Users").Where("email = ?", users.Email).First(users).Error
	if err != nil {
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
	fmt.Println(user_id)

	// Check if user exists
	if err := s.db.Table("Users").Where("user_id = ?", user_id).Error; err != nil {

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
