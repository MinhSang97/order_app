package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/plutov/paypal/v4"
)

func main() {
	// Initialize PayPal client
	client, err := paypal.NewClient("AaTwQs5lkv1030R3J0TNDAEqrsqpk5-obnTlzKgLfYgtt9_rmivhUuvGO-PpKzonTgzNFQ5u1QgTFInf", "EPa-78ktJlm8z0FWnfbZPHWjWQm51yJr33mS0dBMzvlIPAjT8ACiSUg3qlWGTL6uwlmwg0UtefKjbCcL", paypal.APIBaseSandBox)
	if err != nil {
		log.Fatalf("Error initializing PayPal client: %v", err)
	}
	client.SetLog(os.Stdout) // Set log to terminal stdout

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/order/:id", func(c *gin.Context) {
		orderID := c.Param("id")
		order, err := client.GetOrder(context.Background(), orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, order)
	})

	router.POST("/order", func(c *gin.Context) {
		var req struct {
			Total    string `json:"total"`
			Currency string `json:"currency"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		purchaseUnits := []paypal.PurchaseUnitRequest{{
			ReferenceID: "ref-id",
			Amount: &paypal.PurchaseUnitAmount{
				Value:    req.Total,
				Currency: req.Currency,
			},
		}}
		order, err := client.CreateOrder(
			context.Background(),
			paypal.OrderIntentCapture,
			purchaseUnits,
			nil,
			nil,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, order)
	})

	// Confirm order
	router.POST("/order/:id/confirm", func(c *gin.Context) {
		orderID := c.Param("id")
		// Perform the logic to confirm the order here
		// For demonstration purposes, let's just return a success message
		c.JSON(http.StatusOK, gin.H{"message": "Order confirmed successfully", "orderID": orderID})
	})

	// Cancel order
	router.POST("/order/:id/cancel", func(c *gin.Context) {
		orderID := c.Param("id")
		// Perform the logic to cancel the order here
		// For demonstration purposes, let's just return a success message
		c.JSON(http.StatusOK, gin.H{"message": "Order cancelled successfully", "orderID": orderID})
	})

	// Start the server
	router.Run(":8080")
}
