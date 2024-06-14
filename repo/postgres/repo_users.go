package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/model/users_model"
	"github.com/MinhSang97/order_app/pkg/error"
	"github.com/MinhSang97/order_app/pkg/log"
	"github.com/MinhSang97/order_app/repo"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type usersRepository struct {
	db *gorm.DB
}

func (s usersRepository) CreateUsers(ctx context.Context, users *users_model.Users) error {
	// Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return errors.SignUpFail
	}

	expired_at := time.Now().Add(time.Hour * 8)

	// Insert into the users table
	queryUser := `INSERT INTO users (user_id, email, password, name, phone_number, role, created_at, address, telegram, expired_at) VALUES($1, $2, $3, $4, $5, $6,$7,$8,$9,$10);`
	if err := tx.Exec(queryUser, users.UserId, users.Email, users.PassWord, users.Name, users.PhoneNumber, users.Role, time.Now(), users.Address, users.Telegram, expired_at).Error; err != nil {
		tx.Rollback()
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.UserConflict
			}
		}
		return errors.UserConflict
	}

	// Insert into the user_addresses table
	queryUserAddress := `INSERT INTO user_addresses (user_id, address, lat, long, ward_id, ward_text, district_id, district_text, province_id, province_text, national_id, national_text, address_default, name, phone_number) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,$15);`
	if err := tx.Exec(queryUserAddress, users.UserId, users.Address, users.Lat, users.Long, users.WardId, users.WardText, users.DistrictId, users.DistrictText, users.ProvinceId, users.ProvinceText, users.NationalId, users.NationalText, "yes", users.Name, users.PhoneNumber).Error; err != nil {
		tx.Rollback()
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "42P01" {
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

func (s usersRepository) GetUsers(ctx context.Context, users *users_model.ReqUsersSignIn) (*users_model.ReqUsersSignIn, error) {
	token := users.Token

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

	// Step 2: Check if the role is "users"
	if role == "" {
		return nil, errors.UserNotFound
	} else if role != "users" {
		return nil, errors.NotUsers
	}

	err = s.db.Table("users").Where("email = ? OR phone_number = ? AND password = ?", users.Email, users.PhoneNumber, users.PassWord).First(users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserNotFound
		}
		log.Error(err.Error())
		return nil, err
	}

	// Step 3: Proceed with the rest of the function if the role is "users"
	var user_id string
	err = s.db.Table("users").Select("user_id").Where("email = ?", users.Email).Scan(&user_id).Error
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
		users = &users_model.ReqUsersSignIn{
			Email:    users.Email,
			PassWord: users.PassWord,
			Token:    token,
		}
		return users, nil
	} else {
		users = &users_model.ReqUsersSignIn{
			Email:    users.Email,
			PassWord: users.PassWord,
			Token:    token_db.String,
		}
		return users, nil
	}
	return users, nil
}

func (s usersRepository) UpdateUsers(ctx context.Context, user_id string, users *users_model.Users) error {
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

func (s usersRepository) DeleteUsers(ctx context.Context, user_id string) error {
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

	query := `DELETE FROM user_addresses WHERE user_id = ?;`
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

func (s usersRepository) GetAddressUsersFunction(ctx context.Context, user_id string) (*users_model.UsersAddressModel, error) {
	var userAddress users_model.UsersAddressModel

	err := s.db.Table("user_addresses").Where("user_id = ?", user_id).Scan(&userAddress).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.UserAddressNotFound
		}
		return nil, err
	}

	return &userAddress, nil
}

func (s usersRepository) AddAddressUsersFunction(ctx context.Context, user_id string, address *users_model.UsersAddressModel) error {
	err := s.db.Table("user_addresses").Where("user_id = ?", user_id).Update("address_default", "no").Error
	if err != nil {
		return errors.DefaultAddressFail

	}
	// Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return errors.AddAddressFail
	}
	var address_default string
	err = s.db.Table("user_addresses").Select("address_default").Where("user_id = ?", user_id).Scan(&address_default).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.UserAddressNotFound
		}
		return errors.AddAddressFail
	}
	if address_default == "yes" {
		address.AddressDefault = "no"
	}
	// Insert into the user_addresses table
	queryUserAddress := `INSERT INTO user_addresses (user_id, address, lat, long, ward_id, ward_text, district_id, district_text, province_id, province_text, national_id, national_text, address_default, name, phone_number,type_address) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,$15,$16);`
	if err := tx.Exec(queryUserAddress, user_id, address.Address, address.Lat, address.Long, address.WardId, address.WardText, address.DistrictId, address.DistrictText, address.ProvinceId, address.ProvinceText, address.NationalId, address.NationalText, address.AddressDefault, address.Name, address.PhoneNumber, address.TypeAddress).Error; err != nil {
		tx.Rollback()
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.AddAddressFail
			}
		}
		return errors.AddAddressFail
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.AddAddressFail
	}

	return nil

}

func (s usersRepository) DefaultAddressUsersFunction(ctx context.Context, user_id string, address *users_model.UsersAddressModel) error {
	var countUser int64
	err := s.db.Table("users").Where("user_id = ?", user_id).Count(&countUser).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return errors.UserNotFound
		}
		return errors.UserNotFound
	}
	var countAddress int64
	err = s.db.Table("user_addresses").Where("user_id = ?", user_id).Count(&countAddress).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return errors.UserNotFound
		}
		return errors.UserNotFound
	}
	if countUser == 0 || countAddress == 0 {
		return errors.UserAddressNotFound
	}

	err = s.db.Table("user_addresses").Where("user_id = ?", user_id).Update("address_default", "no").Error
	if err != nil {
		return errors.DefaultAddressFail
	}

	err = s.db.Table("user_addresses").Where("user_id = ? and address = ?", user_id, address.Address).Update("address_default", address.AddressDefault).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return errors.DefaultAddressFail
			}
		}
	}
	return nil

}

func (s usersRepository) AddOrderUsersOrder(ctx context.Context, user_id string, order *model.OrderModel) (*model.OrderModel, error) {
	queryOrder := `INSERT INTO orders (order_id, user_id, order_date, total_price, status, address, payment_method) VALUES ($1, $2, $3, $4, $5, $6, $7);`
	if err := s.db.Exec(queryOrder, order.OrderID, user_id, order.OrderDate, order.TotalPrice, order.Status, order.Address, order.PaymentMethod).Error; err != nil {
		fmt.Println(err)
		return nil, errors.AddOrderFail
	}

	//Insert into the payments table
	queryPayment := `INSERT INTO payments (order_id, payment_date, amount, payment_method) VALUES ($1, $2, $3, $4);`
	if err := s.db.Exec(queryPayment, order.OrderID, order.PaymentDate, order.Amount, order.PaymentMethod).Error; err != nil {
		return nil, errors.AddPaymentFail
	}

	//Insert into the order_discounts table
	queryOrderDiscounts := `INSERT INTO order_discounts (order_id, discount_code_id) VALUES ($1, $2);`
	if err := s.db.Exec(queryOrderDiscounts, order.OrderID, order.DiscountCodeId).Error; err != nil {
		return nil, errors.AddOrderDiscountsFail

	}

	//Insert into the history_transaction table
	queryHistoryTransaction := `INSERT INTO history_transaction (order_id, user_id, amount) VALUES ($1, $2, $3);`
	if err := s.db.Exec(queryHistoryTransaction, order.OrderID, user_id, order.Amount).Error; err != nil {
		return nil, errors.AddHistoryTransactionFail
	}

	//Insert into the order_items table
	queryOrderItems := `INSERT INTO order_items (order_id, item_id, quantity, price) VALUES($1, $2, $3, $4);`
	for i, item := range order.ItemID {
		if err := s.db.Exec(queryOrderItems, order.OrderID, item, order.Quantity[i], order.Price[i]).Error; err != nil {
			return nil, errors.AddOrderItemsFail
		}

	}
	return order, nil
}

func (s usersRepository) StatusOrderUserOrder(ctx context.Context, user_id string, order *model.OrderModel) error {
	query := `UPDATE orders SET status = ? WHERE order_id = ? and user_id = ?;`
	if err := s.db.Exec(query, order.Status, order.OrderID).Error; err != nil {
		return errors.StatusOrderFail
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
