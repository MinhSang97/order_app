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
	expired_at := time.Now().Add(time.Hour * 8)

	// Insert into the users table
	addUsers := `INSERT INTO users (user_id, password, name, email, phone_number, address, role, created_at, expired_at) 
					VALUES($1,$2,$3,$4,$5,$6,$7,$8, $9);`
	err := tx.Exec(addUsers, users.UserId, users.Password, users.Name, users.Email, users.PhoneNumber, users.Address, users.Role, time.Now(), expired_at).Error
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
	query := `INSERT INTO user_addresses (user_id, address, name, phone_number) VALUES($1, $2, $3, $4);`
	if err := tx.Exec(query, users.UserId, users.Address, users.Name, users.PhoneNumber).Error; err != nil {
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
	type MenuItemWithCustomization struct {
		ItemID              string
		ItemName            string
		Description         string
		Price               float64
		ImageUrl            string
		CustomizationOption string
		ExtraPrice          float64
	}

	// Custom SQL query to join the menu_items and item_customizations tables
	query := `
	SELECT a.item_id, a.item_name, a.description, a.price, a.image_url, b.customization_option, b.extra_price 
    FROM menu_items a
    LEFT JOIN item_customizations b
    ON a.item_id = b.item_id;`

	var results []MenuItemWithCustomization

	if err := s.db.Raw(query).Scan(&results).Error; err != nil {
		return nil, fmt.Errorf("get all menu items error: %w", err)
	}

	// Create a map to aggregate customization options and extra prices
	menuMap := make(map[string]*model.MenuItemsModel)

	for _, result := range results {
		if item, exists := menuMap[result.ItemID]; exists {
			item.CustomizationOption = append(item.CustomizationOption, result.CustomizationOption)
			item.ExtraPrice = append(item.ExtraPrice, result.ExtraPrice)
		} else {
			menuMap[result.ItemID] = &model.MenuItemsModel{
				ItemID:              result.ItemID,
				ItemName:            result.ItemName,
				Description:         result.Description,
				Price:               result.Price,
				ImageUrl:            result.ImageUrl,
				CustomizationOption: []string{result.CustomizationOption},
				ExtraPrice:          []float64{result.ExtraPrice},
			}
		}
	}

	// Convert the map to a slice
	for _, item := range menuMap {
		menu = append(menu, *item)
	}

	return menu, nil
}
func (s adminFunctionRepository) AddMenu(ctx context.Context, menu *model.MenuItemsModel) (*model.MenuItemsModel, error) {
	query := `INSERT INTO menu_items (item_id, item_name, description, price, image_url) VALUES($1, $2, $3, $4, $5) RETURNING item_id;`
	err := s.db.Raw(query, menu.ItemID, menu.ItemName, menu.Description, menu.Price, menu.ImageUrl).Scan(&menu).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return nil, errors.AddMenuItemsFail
			}
		}
		return nil, errors.AddMenuItemsFail
	}

	query_item_customizations := `INSERT INTO item_customizations (item_id, customization_option, extra_price) VALUES($1, $2, $3);`
	for i := 0; i < len(menu.CustomizationOption); i++ {
		err = s.db.Exec(query_item_customizations, menu.ItemID, menu.CustomizationOption[i], menu.ExtraPrice[i]).Error
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok {
				if pgErr.Code == "22P02" {
					return nil, errors.AddMenuItemsFail
				}
			}
			return nil, errors.AddMenuItemsFail
		}
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

	query_menu_items := `UPDATE menu_items SET item_name = $1, description = $2, price = $3, image_url = $4 WHERE item_id = $5;`
	err = s.db.Exec(query_menu_items, menu.ItemName, menu.Description, menu.Price, menu.ImageUrl, item_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.EditMenuItemsFail
			}
		}
		return errors.EditMenuItemsFail

	}

	err = s.db.Exec("DELETE FROM item_customizations WHERE item_id = $1", item_id).Error
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "22P02" {
				return errors.EditMenuItemsFail
			}
		}
		return errors.EditMenuItemsFail
	}

	query_item_customizations := `INSERT INTO item_customizations (item_id, customization_option, extra_price) VALUES($1, $2, $3);`
	for i := 0; i < len(menu.CustomizationOption); i++ {
		err = s.db.Exec(query_item_customizations, item_id, menu.CustomizationOption[i], menu.ExtraPrice[i]).Error
		if err != nil {
			if pgErr, ok := err.(*pq.Error); ok {
				if pgErr.Code == "22P02" {
					return errors.AddMenuItemsFail
				}
			}
			return errors.AddMenuItemsFail
		}
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

// admin_function_feedback
func (s adminFunctionRepository) GetFeedbackAll(ctx context.Context) ([]model.FeedbackModel, error) {
	var feedback []model.FeedbackModel
	if err := s.db.Table("feedbacks").Scan(&feedback).Error; err != nil {
		return feedback, fmt.Errorf("get all feedback error: %w", err)
	}
	return feedback, nil
}

var instance adminFunctionRepository

func NewAdminFunctionUseCase(db *gorm.DB) repo.AdminFunctionRepo {
	if instance.db == nil {
		instance.db = db
	}
	return instance
}
