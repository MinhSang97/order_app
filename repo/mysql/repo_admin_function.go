package mysql

import (
	"context"
	"encoding/json"
	"fmt"
	errors "github.com/MinhSang97/order_app/error"
	"github.com/MinhSang97/order_app/model/admin_model"
	"github.com/MinhSang97/order_app/redis"
	"github.com/MinhSang97/order_app/repo"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type adminFunctionRepository struct {
	db *gorm.DB
}

var RedisClient = redis.ConnectRedis()

func (s adminFunctionRepository) AddUser(ctx context.Context, users *admin_model.AdminFunctionModel) error {
	// Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return errors.SignUpFail
	}

	// Insert into the users table
	addUsers := `INSERT INTO order_app.users (user_id, pass_word, name, email, phone_number, address, role, created_at) 
					VALUES(?,?,?,?,?,?,?,?);`
	err := tx.Exec(addUsers, users.UserId, users.Password, users.Name, users.Name, users.PhoneNumber, users.Address, users.Role, time.Now()).Error
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
	if err := tx.Exec(query, users.UserId, users.Address).Error; err != nil {
		tx.Rollback()
		return errors.SignUpFail
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.SignUpFail
	}

	return nil
}

func (s adminFunctionRepository) GetAll(ctx context.Context) ([]admin_model.AdminFunctionModel, error) {
	var users []admin_model.AdminFunctionModel

	// Đọc danh sách users từ Redis (nếu có)
	cachedUsersJSON, err := RedisClient.Get(ctx, "users").Result()
	if err == nil {
		var cachedUserss []admin_model.AdminFunctionModel
		err := json.Unmarshal([]byte(cachedUsersJSON), &cachedUserss)
		if err != nil {
			log.Println("Failed to unmarshal users from Redis:", err)
			return users, fmt.Errorf("Failed to unmarshal users from Redis: %w", err)
		}
		fmt.Println("Users fetched from Redis")
		return cachedUserss, nil
	}

	// Nếu không tìm thấy trong Redis, đọc từ cơ sở dữ liệu MySQL
	if err := s.db.Table("users").Scan(&users).Error; err != nil {
		return users, fmt.Errorf("get all users error: %w", err)
	}
	fmt.Println("Users", users)

	// Cache danh sách sinh viên vào Redis
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Failed to marshal students:", err)
		return users, fmt.Errorf("Failed to marshal users: %w", err)
	}

	err = redis.RedisClient.Set(ctx, "users", jsonUsers, 0).Err()
	if err != nil {
		fmt.Println("Failed to cache users in Redis:", err)
	}

	fmt.Println("Students query from MySQL")

	return users, nil
}

func (s adminFunctionRepository) Edit(ctx context.Context, user_id string, users *admin_model.AdminFunctionModel) error {

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

var instance adminFunctionRepository

func NewAdminFunctionUseCase(db *gorm.DB) repo.AdminFunctionRepo {
	if instance.db == nil {
		instance.db = db
	}
	return instance
}
