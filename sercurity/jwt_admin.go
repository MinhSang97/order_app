//package sercurity
//
//import (
//	"fmt"
//	"github.com/MinhSang97/order_app/usecases/dto/admin_dto"
//
//	"github.com/golang-jwt/jwt"
//	"time"
//)
//
//const SECRET_KEY_ADMIN = "learngolanglalalafdfds"
//
//type JwtCustomClaims struct {
//	UserId string
//	Role   string
//	jwt.StandardClaims
//}
//
//func GenTokenAdmin(user admin_dto.Admin) (string, error) {
//	claims := &JwtCustomClaims{
//		UserId: user.UserId,
//		Role:   user.Role,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	result, err := token.SignedString([]byte(SECRET_KEY_ADMIN))
//	if err != nil {
//		fmt.Println("Loi tao token", err.Error())
//		return "Tao Token Loi!", err
//	}
//
//	return result, nil
//
//}

package sercurity

import (
	"fmt"
	"github.com/MinhSang97/order_app/usecases/dto/admin_dto"
	"github.com/golang-jwt/jwt"
	"sync"
	"time"
)

const SECRET_KEY_ADMIN = "learngolanglalalafdfds"

type JwtCustomClaims struct {
	UserId string
	Role   string
	jwt.StandardClaims
}

var tokenStore = struct {
	sync.RWMutex
	tokens map[string]string
	expiry map[string]int64
}{tokens: make(map[string]string), expiry: make(map[string]int64)}

func GenTokenAdmin(user admin_dto.Admin) (string, error) {
	tokenStore.RLock()
	existingToken, exists := tokenStore.tokens[user.UserId]
	expiryTime, _ := tokenStore.expiry[user.UserId]
	tokenStore.RUnlock()

	if exists && time.Now().Unix() < expiryTime {
		return existingToken, nil
	}

	claims := &JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY_ADMIN))
	if err != nil {
		fmt.Println("Lỗi tạo token:", err.Error())
		return "Tạo Token Lỗi!", err
	}

	tokenStore.Lock()
	tokenStore.tokens[user.UserId] = result
	tokenStore.expiry[user.UserId] = claims.ExpiresAt
	tokenStore.Unlock()

	return result, nil
}
