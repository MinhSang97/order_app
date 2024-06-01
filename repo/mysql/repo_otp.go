package mysql

import (
	"context"
	"fmt"
	errors "github.com/MinhSang97/order_app/error"
	"github.com/MinhSang97/order_app/log"
	"github.com/MinhSang97/order_app/model"
	"github.com/MinhSang97/order_app/repo"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

type otpRepository struct {
	db *gorm.DB
}

func (s otpRepository) SendOtp(ctx context.Context, otp *model.OtpModel) error {
	// Start a transaction
	var user_id string
	var cN int64
	err := s.db.Table("Users").Select("user_id").Where("email = ?", otp.Email).Scan(&user_id).Count(&cN).Error
	if err != nil {
		log.Error(err.Error())
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user not found with email: %s", otp.Email)
		}
		return err
	}
	if cN == 0 {
		return fmt.Errorf("user not found with email: %s", otp.Email)
	} else {
		otpSave := model.OtpModel{
			UserId:    user_id,
			Email:     otp.Email,
			Otp:       otp.Otp,
			CreatedAt: time.Now(),
		}
		var count int64
		err := s.db.Table("recover_password").Where("user_id = ? AND email = ?", user_id, otpSave.Email).Count(&count).Error
		if err != nil {
			return err
		}
		//fmt.Println("Số lượng bản ghi:", checkOTP)
		if count == 1 {
			// Tìm thấy bản ghi
			query := `UPDATE order_app.recover_password SET otp= ?, created_at= ? WHERE user_id= ? AND email=?;`
			if err := s.db.Exec(query, otpSave.Otp, otpSave.CreatedAt, otpSave.UserId, otpSave.Email).Error; err != nil {
				if driverErr, ok := err.(*mysql.MySQLError); ok {
					if driverErr.Number == 1452 {
						return errors.CreatOTPFail
					}
				}
				return errors.CreatOTPFail
			}

		} else {
			// Không tìm thấy bản ghi
			query := `INSERT INTO order_app.recover_password ( user_id, email, otp, created_at) VALUES( ?, ?, ?, ?);`
			if err := s.db.Exec(query, otpSave.UserId, otpSave.Email, otpSave.Otp, otpSave.CreatedAt).Error; err != nil {
				if driverErr, ok := err.(*mysql.MySQLError); ok {
					if driverErr.Number == 1452 {
						return errors.CreatOTPFail
					}
				}
				return errors.CreatOTPFail
			}
		}
	}
	return nil
}

func (s otpRepository) VerifyOtp(ctx context.Context, otp *model.OtpModel) error {
	var count int64
	err := s.db.Table("recover_password").Where("email = ? AND otp = ?", otp.Email, otp.Otp).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.OTPVerified
	} else {

	}
	return nil
}

func (s otpRepository) ChangePassword(ctx context.Context, otp_code string, otp *model.OtpModel) error {
	fmt.Println(otp_code, otp)
	var user_id string
	var count int64

	err := s.db.Table("recover_password").Where("email = ? AND otp = ?", otp.Email, otp_code).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.OTPVerified
	} else {
		err := s.db.Table("Users").Select("user_id").Where("email = ?", otp.Email).Scan(&user_id).Error
		if err != nil {
			log.Error(err.Error())
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("user not found with email: %s", otp.Email)
			}
			return err
		}
		fmt.Println(user_id)
		query_users := `UPDATE order_app.users SET pass_word= ? WHERE email= ? AND user_id = ?;`
		if err := s.db.Exec(query_users, otp.PassWordNew, otp.Email, user_id).Error; err != nil {
			if driverErr, ok := err.(*mysql.MySQLError); ok {
				if driverErr.Number == 1452 {
					return errors.ChangePasswordByOTPFail
				}
			}
			return errors.ChangePasswordByOTPFail
		}
		query_recover_password := `UPDATE order_app.recover_password SET password_new= ? WHERE otp= ? AND email=? AND user_id = ?;`
		if err := s.db.Exec(query_recover_password, otp.PassWordNew, otp_code, otp.Email, user_id).Error; err != nil {
			if driverErr, ok := err.(*mysql.MySQLError); ok {
				if driverErr.Number == 1452 {
					return errors.ChangePasswordByOTPFail
				}
			}
			return errors.ChangePasswordByOTPFail
		}
	}

	return nil
}

var instancess otpRepository

func NewOtpRepository(db *gorm.DB) repo.OtpRepo {
	if instancess.db == nil {
		instancess.db = db

	}
	return instancess
}
