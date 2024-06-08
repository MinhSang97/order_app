package main

import (
	"github.com/MinhSang97/order_app/redis"
	"github.com/MinhSang97/order_app/router"
)

func main() {
	router.Route()
	redis.ConnectRedis()

}
