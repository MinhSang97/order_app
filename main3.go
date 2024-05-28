package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var otpStore = make(map[string]string)

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

func sendEmail(to string, otp string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is: %s", otp))

	d := gomail.NewDialer("smtp.gmail.com", 587, "minhsangnguyen463@gmail.com", "dcdb cqtj dfsh izwv")

	if err := d.DialAndSend(m); err != nil {
		// Log the error
		log.Printf("Failed to send email to %s: %v", to, err)
		return err
	}
	return nil
}

func sendOTP(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	otp := generateOTP()
	otpStore[email] = otp

	if err := sendEmail(email, otp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email"})
}

func verifyOTP(c *gin.Context) {
	email := c.Query("email")
	otp := c.Query("otp")

	if email == "" || otp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and OTP are required"})
		return
	}

	storedOTP, exists := otpStore[email]
	if !exists || storedOTP != otp {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	delete(otpStore, email)
	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}

func main() {
	router := gin.Default()

	router.GET("/send-otp", sendOTP)
	router.GET("/verify-otp", verifyOTP)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
