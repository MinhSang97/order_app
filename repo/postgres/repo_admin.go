package postgres

import (
	"context"
	"database/sql"
	"github.com/MinhSang97/order_app/model/admin_model"
	"github.com/MinhSang97/order_app/model/users_model"
	"github.com/MinhSang97/order_app/pkg/error"
	"github.com/MinhSang97/order_app/pkg/log"
	"github.com/MinhSang97/order_app/repo"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type adminRepository struct {
	db *gorm.DB
}

func (s adminRepository) CreateAdmin(ctx context.Context, admin *admin_model.Admin) error {

	//Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return errors.SignUpFail
	}

	expired_at := time.Now().Add(time.Hour * 8)

	// Insert into the users table
	queryUser := `INSERT INTO users (user_id, email, password, name, phone_number, role, created_at, address, telegram, expired_at) VALUES($1, $2, $3, $4, $5, $6,$7,$8,$9, $10);`
	if err := tx.Exec(queryUser, admin.UserId, admin.Email, admin.PassWord, admin.Name, admin.PhoneNumber, admin.Role, time.Now(), admin.Address, admin.Telegram, expired_at).Error; err != nil {
		tx.Rollback()
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.UserConflict
			}
		}
		return errors.UserConflict
	}

	// Insert into the user_addresses table
	queryUserAddress := `INSERT INTO user_addresses (user_id, address, lat, long, ward_id, ward_text, district_id, district_text, province_id, province_text, national_id, national_text, address_default, name, phone_number) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15);`
	if err := tx.Exec(queryUserAddress, admin.UserId, admin.Address, admin.Lat, admin.Long, admin.WardId, admin.WardText, admin.DistrictId, admin.DistrictText, admin.ProvinceId, admin.ProvinceText, admin.NationalId, admin.NationalText, "no", admin.Name, admin.PhoneNumber).Error; err != nil {
		tx.Rollback()
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.SignUpFail
			}

		}
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
	token := admin.Token

	// Step 1: Query the database to get the user's role based on their email
	var role string
	err := s.db.Table("users").Select("role").Where("email = ?", users.Email).Scan(&role).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		return nil, err
	}

	// Step 2: Check if the role is "admin"
	if role == "" {
		return nil, errors.UserNotFound
	} else if role != "admin" {
		return nil, errors.NotAdmin
	}

	err = s.db.Table("users").Where("email = ? or phone_number = ? ", users.Email, users.PhoneNumber).First(users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		log.Error(err.Error())
		return nil, err
	}

	// Step 3: Proceed with the rest of the function if the role is "admin"
	var user_id string
	err = s.db.Table("users").Select("user_id").Where("phone_number = ?", users.PhoneNumber).Scan(&user_id).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		return nil, err
	}

	var expired_at time.Time
	err = s.db.Table("users").Select("expired_at").Where("user_id = ?", user_id).Scan(&expired_at).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		return nil, err
	}

	// Truncate to remove milliseconds and timezone info
	expired_at_str := expired_at.Format("2006-01-02 15:04:05")
	expired_at, err = time.Parse("2006-01-02 15:04:05", expired_at_str)
	if err != nil {
		return nil, err
	}

	// Get current time in UTC and truncate to remove milliseconds
	expired_now := time.Now()
	expired_now_str := expired_now.Format("2006-01-02 15:04:05")
	expired_now, err = time.Parse("2006-01-02 15:04:05", expired_now_str)
	if err != nil {
		return nil, err
	}

	kq := expired_now.Sub(expired_at)

	var token_db sql.NullString
	err = s.db.Table("users").Select("token").Where("user_id = ?", user_id).Scan(&token_db).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		return nil, err
	}

	expired_at_db := time.Now().Add(time.Hour * 8)

	if kq > 8*time.Hour || token_db.String == "" {
		queryInsertToken := `UPDATE users SET token = ?, expired_at = ? WHERE user_id = ?;`
		if err := s.db.Exec(queryInsertToken, token, expired_at_db, user_id).Error; err != nil {
			return nil, errors.SignInFail
		}
	}
	if kq > 8*time.Hour || token_db.String == "" {
		users = &admin_model.ReqSignIn{
			Email:    users.Email,
			PassWord: users.PassWord,
			Token:    token,
			UserID:   user_id,
		}
		return users, nil
	} else {
		users = &admin_model.ReqSignIn{
			Email:    users.Email,
			PassWord: users.PassWord,
			Token:    token_db.String,
			UserID:   user_id,
		}
		return users, nil
	}
	return users, nil
}

func (s adminRepository) UpdateAdmin(ctx context.Context, user_id string, admin *admin_model.Admin) error {
	users := admin

	queryUpdate := `UPDATE users SET name = ?, sex = ?, birth_date = ?, telegram = ?, updated_at = ? WHERE user_id = ?;`
	if err := s.db.Exec(queryUpdate, users.Name, users.Sex, users.BirthDate, users.Telegram, time.Now(), user_id).Error; err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "42P01" {
				return errors.UserNotUpdated
			}
		}
		return errors.UserNotUpdated
	}
	return nil
}

func (s adminRepository) DeleteAdmin(ctx context.Context, user_id string) error {
	var user users_model.Users

	// Check if user exists
	if err := s.db.Table("users").Where("user_id = ?", user_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.UserNotFound
		}
		return err
	}
	if user_id == "" {
		return errors.UserNotFound
	}

	query := `DELETE FROM user_addresses WHERE user_id = $1;`
	err := s.db.Exec(query, user_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.UserNotDeleted
			}
		}
		return err
	}

	// If user exists, delete the user
	if err := s.db.Table("users").Where("user_id = ?", user_id).Delete(&user).Error; err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
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
