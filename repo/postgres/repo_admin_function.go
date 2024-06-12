package postgres

import (
	"context"
	"fmt"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/model/admin_model"
	"github.com/MinhSang97/order_app/pkg/error"
	"github.com/MinhSang97/order_app/pkg/log"
	"github.com/MinhSang97/order_app/redis"
	"github.com/MinhSang97/order_app/repo"
	"github.com/lib/pq"
	"gorm.io/gorm"
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
	addUsers := `INSERT INTO users (user_id, pass_word, name, email, phone_number, address, role, created_at) 
					VALUES($1,$2,$3,$4,$5,$6,$7,$8);`
	err := tx.Exec(addUsers, users.UserId, users.Password, users.Name, users.Email, users.PhoneNumber, users.Address, users.Role, time.Now()).Error
	if err != nil {
		tx.Rollback()
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23514" {
				return errors.UserConflict
			}
		}
		return errors.SignUpFail
	}

	// Insert into the user_addresses table
	query := `INSERT INTO user_addresses (user_id, address) VALUES($1, $2);`
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

	//// Đọc danh sách users từ Redis (nếu có)
	//cachedUsersJSON, err := RedisClient.Get(ctx, "users").Result()
	//if err == nil {
	//	var cachedUserss []admin_model.AdminFunctionModel
	//	err := json.Unmarshal([]byte(cachedUsersJSON), &cachedUserss)
	//	if err != nil {
	//		log.Println("Failed to unmarshal users from Redis:", err)
	//		return users, fmt.Errorf("Failed to unmarshal users from Redis: %w", err)
	//	}
	//	fmt.Println("Users fetched from Redis")
	//	return cachedUserss, nil
	//}

	// Nếu không tìm thấy trong Redis, đọc từ cơ sở dữ liệu MySQL
	if err := s.db.Table("users").Scan(&users).Error; err != nil {
		return users, fmt.Errorf("get all users error: %w", err)
	}

	//// Cache danh sách sinh viên vào Redis
	//jsonUsers, err := json.Marshal(users)
	//if err != nil {
	//	fmt.Println("Failed to marshal students:", err)
	//	return users, fmt.Errorf("Failed to marshal users: %w", err)
	//}
	//
	//err = redis.RedisClient.Set(ctx, "users", jsonUsers, 0).Err()
	//if err != nil {
	//	fmt.Println("Failed to cache users in Redis:", err)
	//}
	//
	//fmt.Println("users query from Postgres")

	return users, nil
}

func (s adminFunctionRepository) Edit(ctx context.Context, user_id string, users *admin_model.AdminFunctionModel) error {
	err := s.db.Table("users").Where("user_id = ?", user_id).Updates(users).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "42P01" {
				return errors.UserNotUpdated
			}
		}
		return errors.SignUpFail
	}
	return nil
}

func (s adminFunctionRepository) DeleteUsers(ctx context.Context, email string) error {
	var user_id string
	// Check if user exists
	if err := s.db.Table("users").Where("email = ?", email).Select("user_id").Scan(&user_id).Error; err != nil {
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
			if pgErr.Code == "42P01" {
				return errors.UserNotDeleted
			}
		}
		return err
	}

	// If user exists, delete the user
	deleteUsers := `DELETE FROM users WHERE user_id = $1;`
	if err := s.db.Exec(deleteUsers, user_id).Error; err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "42P01" {
				return errors.UserNotDeleted
			}
		}
		return err
	}

	return nil
}

// admin_function_menu

func (s adminFunctionRepository) GetMenuAll(ctx context.Context) ([]model.MenuItemsModel, error) {
	var menu []model.MenuItemsModel
	if err := s.db.Table("menu_items").Scan(&menu).Error; err != nil {
		return menu, fmt.Errorf("get all menu items error: %w", err)
	}
	return menu, nil
}

func (s adminFunctionRepository) AddMenu(ctx context.Context, menu *model.MenuItemsModel) (*model.MenuItemsModel, error) {
	query := `INSERT INTO menu_items (item_id, name, description, price, image_url) VALUES($1, $2, $3, $4, $5) RETURNING item_id;`
	err := s.db.Raw(query, menu.ItemID, menu.Name, menu.Description, menu.Price, menu.ImageUrl).Scan(&menu).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return nil, errors.AddMenuItemsFail
			}
		}
		return nil, errors.AddMenuItemsFail
	}

	fmt.Println(menu.ExtraPrice1)

	query_item_customizations := `INSERT INTO item_customizations (item_id, customization_option_1, extra_price_1, customization_option_2, extra_price_2, customization_option_3, extra_price_3, customization_option_4, extra_price_4, customization_option_5, extra_price_5, customization_option_6, extra_price_6, customization_option_7, extra_price_7, customization_option_8, extra_price_8, customization_option_9, extra_price_9, customization_option_10, extra_price_10)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21);`
	err = s.db.Exec(query_item_customizations, menu.ItemID, menu.CustomizationOption1, menu.ExtraPrice1, menu.CustomizationOption2, menu.ExtraPrice2, menu.CustomizationOption3, menu.ExtraPrice3, menu.CustomizationOption4, menu.ExtraPrice4, menu.CustomizationOption5, menu.ExtraPrice5, menu.CustomizationOption6, menu.ExtraPrice6, menu.CustomizationOption7, menu.ExtraPrice7, menu.CustomizationOption8, menu.ExtraPrice8, menu.CustomizationOption9, menu.ExtraPrice9, menu.CustomizationOption10, menu.ExtraPrice10).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return nil, errors.AddMenuItemsFail
			}
		}
		return nil, errors.AddMenuItemsFail
	}

	return menu, nil
}

func (s adminFunctionRepository) EditMenu(ctx context.Context, item_id string, menu *model.MenuItemsModel) error {
	var count int64
	err := s.db.Table("menu_items").Where("item_id = ?", item_id).Count(&count).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return errors.ItemIDNotFound
		}
		return errors.ItemIDNotFound
	}
	if count == 0 {
		return errors.ItemIDNotFound

	}

	query_menu_items := `UPDATE menu_items SET name = $1, description = $2, price = $3, image_url = $4 WHERE item_id = $5;`
	err = s.db.Exec(query_menu_items, menu.Name, menu.Description, menu.Price, menu.ImageUrl, item_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.EditMenuItemsFail
			}
		}
		return errors.EditMenuItemsFail

	}

	query_item_customizations := `UPDATE item_customizations SET customization_option_1 = $1, extra_price_1 = $2, customization_option_2 = $3, extra_price_2 = $4, customization_option_3 = $5, extra_price_3 = $6, customization_option_4 = $7, extra_price_4 = $8, customization_option_5 = $9, extra_price_5 = $10, customization_option_6 = $11, extra_price_6 = $12, customization_option_7 = $13, extra_price_7 = $14, customization_option_8 = $15, extra_price_8 = $16, customization_option_9 = $17, extra_price_9 = $18, customization_option_10 = $19, extra_price_10 = $20 WHERE item_id = $21;`
	err = s.db.Exec(query_item_customizations, menu.CustomizationOption1, menu.ExtraPrice1, menu.CustomizationOption2, menu.ExtraPrice2, menu.CustomizationOption3, menu.ExtraPrice3, menu.CustomizationOption4, menu.ExtraPrice4, menu.CustomizationOption5, menu.ExtraPrice5, menu.CustomizationOption6, menu.ExtraPrice6, menu.CustomizationOption7, menu.ExtraPrice7, menu.CustomizationOption8, menu.ExtraPrice8, menu.CustomizationOption9, menu.ExtraPrice9, menu.CustomizationOption10, menu.ExtraPrice10, item_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.EditMenuItemsFail
			}
		}
		return errors.EditMenuItemsFail

	}
	return nil
}

func (s adminFunctionRepository) DeleteMenu(ctx context.Context, item_id string) error {
	var count int64
	err := s.db.Table("menu_items").Where("item_id = ?", item_id).Count(&count).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return errors.ItemIDNotFound
		}
		return errors.ItemIDNotFound
	}
	if count == 0 {
		return errors.ItemIDNotFound
	}

	query := `DELETE FROM item_customizations WHERE item_id = $1;`
	err = s.db.Exec(query, item_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.DeleteMenuItemsFail
			}
		}
		return errors.DeleteMenuItemsFail
	}

	query = `DELETE FROM menu_items WHERE item_id = $1;`
	err = s.db.Exec(query, item_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.DeleteMenuItemsFail
			}
		}
		return errors.DeleteMenuItemsFail
	}
	return nil
}

// admin_function_discount

func (s adminFunctionRepository) GetDiscountAll(ctx context.Context) ([]model.DiscountCodesModel, error) {
	var discount []model.DiscountCodesModel
	if err := s.db.Table("discount_codes").Scan(&discount).Error; err != nil {
		return discount, fmt.Errorf("get all discount codes error: %w", err)
	}
	return discount, nil

}
func (s adminFunctionRepository) AddDiscount(ctx context.Context, discount *model.DiscountCodesModel) (*model.DiscountCodesModel, error) {

	err := s.db.Table("discount_codes").Where("code = ?", discount.Code).First(&discount).Error
	if err == nil {
		return nil, errors.DiscountCodeConflict
	}
	err = s.db.Table("discount_codes").Create(discount).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return nil, errors.AddDiscountFail
			}
		}
		return nil, errors.AddDiscountFail

	}
	return discount, nil
}

func (s adminFunctionRepository) EditDiscount(ctx context.Context, discount_code_id string, discount *model.DiscountCodesModel) error {
	var count int64
	err := s.db.Table("discount_codes").Where("discount_code_id = ?", discount_code_id).Count(&count).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return errors.DiscountCodeIDNotFound
		}
		return errors.DiscountCodeIDNotFound
	}
	if count == 0 {
		return errors.DiscountCodeIDNotFound
	}

	err = s.db.Table("discount_codes").Where("discount_code_id = ?", discount_code_id).Updates(discount).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.EditDiscountFail
			}
		}
		return errors.EditDiscountFail
	}
	return nil
}

func (s adminFunctionRepository) DeleteDiscount(ctx context.Context, discount_code_id string) error {
	var count int64
	err := s.db.Table("discount_codes").Where("discount_code_id = ?", discount_code_id).Count(&count).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return errors.DiscountCodeIDNotFound
		}
		return errors.DiscountCodeIDNotFound
	}
	if count == 0 {
		return errors.DiscountCodeIDNotFound
	}

	query := `DELETE FROM discount_codes WHERE discount_code_id = $1;`
	err = s.db.Exec(query, discount_code_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.DeleteDiscountFail
			}
		}
		return errors.DeleteDiscountFail
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
