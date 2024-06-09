package otp

import (
	"fmt"
	"github.com/MinhSang97/order_app/pkg/log"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendEmail(to string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Printf("Invalid SMTP port: %v", err)
		return err
	}

	d := gomail.NewDialer(smtpHost, port, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		// Log the error
		log.Printf("Failed to send email to %s: %v", to, err)
		return err
	}
	return nil
}
