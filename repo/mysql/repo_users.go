package mysql

import (
	"context"
	errors "github.com/MinhSang97/order_app/error"
	"github.com/MinhSang97/order_app/log"
	"github.com/MinhSang97/order_app/model/users_model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func (s usersRepository) CreateUsers(ctx context.Context, users *users_model.Users) error {
	err := s.db.Table("Users").Create(users).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {

			if driverErr.Number == 1062 {
				return errors.UserConflict
			}
		}
		return errors.SignUpFail
	}

	// Insert into the user_addresses table
	query := `INSERT INTO order_app.user_addresses (user_id, address) VALUES(?, ?);`
	err = s.db.Exec(query, users.UserId, users.Address).Error
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

func (s usersRepository) GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error) {

	// Step 1: Query the database to get the user's role based on their email
	var role string
	err := s.db.Table("Users").Select("role").Where("email = ?", users.Email).Scan(&role).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotAdmin
		}
		return nil, err
	}

	// Step 2: Check if the role is "users"
	if role != "users" {
		return nil, errors.NotUsers
	}

	// Step 3: Proceed with the rest of the function if the role is "users"
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

func (s usersRepository) UpdateUsers(ctx context.Context, user_id string, users *users_model.Users) error {
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

func (s usersRepository) DeleteUsers(ctx context.Context, user_id string) error {
	var user users_model.Users

	// Check if user exists
	if err := s.db.Table("Users").Where("user_id = ?", user_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.UserNotFound
		}
		return err
	}

	if user_id == "" {
		return errors.UserNotFound
	}

	query := `DELETE FROM order_app.user_addresses WHERE user_id = ?;`
	err := s.db.Exec(query, user_id).Error
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {

			if driverErr.Number == 1062 {
				return errors.UserNotDeleted
			}
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

var instancesUsers usersRepository

func NewUsersRepository(db *gorm.DB) repo.UsersRepo {
	if instancesUsers.db == nil {
		instancesUsers.db = db

	}
	return instancesUsers
}
