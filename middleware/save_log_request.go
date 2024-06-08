package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveLogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy ngày hiện tại và định dạng thành chuỗi "2006-01-02"
		currentDate := time.Now().Format("2006-01-02")

		// Tạo đường dẫn thư mục và tên tệp tin log
		logDirectory := "log_middleware"
		logFileName := logDirectory + "/log_middleware_" + currentDate + ".json"

		// Kiểm tra xem thư mục log có tồn tại không, nếu không tạo mới
		if _, err := os.Stat(logDirectory); os.IsNotExist(err) {
			err := os.Mkdir(logDirectory, os.ModePerm)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create log directory"})
				return
			}
		}

		// Mở hoặc tạo một tệp tin để lưu log middleware
		logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to open log file"})
			return
		}
		defer logFile.Close()

		// Ghi log request vào tệp tin đã mở
		log.SetOutput(logFile)
		log.Printf("%s - [%s] \"%s %s %s\"\n",
			c.ClientIP(),
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto)

		// Chuyển quyền kiểm soát sang middleware hoặc handler tiếp theo trong chuỗi middleware
		c.Next()
	}
}
