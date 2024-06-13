package id

import (
	"fmt"
	"strconv"
	"time"
)

func OrderID() int64 {
	// Lấy thời gian hiện tại
	t := time.Now()

	// Chuyển đổi thời gian thành Unix timestamp
	timestamp := t.UnixNano() // Sử dụng UnixNano để có nhiều chữ số hơn

	// Chuyển Unix timestamp thành chuỗi
	order_id_str := fmt.Sprintf("%d", timestamp)

	// Lấy 16 chữ số đầu tiên (hoặc bao nhiêu tùy ý bạn)
	if len(order_id_str) > 16 {
		order_id_str = order_id_str[:16]
	}

	// Chuyển chuỗi thành int64
	order_id, err := strconv.ParseInt(order_id_str, 10, 64)
	if err != nil {
		fmt.Println("Lỗi chuyển đổi:", err)
		return 0
	}
	return order_id
}
