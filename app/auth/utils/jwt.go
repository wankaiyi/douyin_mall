package utils

import (
	"context"
	"douyin_mall/auth/biz/dal/redis"
	"douyin_mall/auth/conf"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var (
	ctx                        = context.Background()
	jwtSecret                  = []byte(conf.GetConf().Jwt.Secret)
	redisClient                = redis.RedisClient
	accessTokenExpireDuration  = time.Hour * 2
	refreshTokenExpireDuration = time.Hour * 24 * 7
)

const (
	// TokenValid 令牌有效
	TokenValid = iota
	// TokenInvalid 令牌不合法
	TokenInvalid
	// TokenExpired 令牌过期
	TokenExpired
)

func GenerateRefreshToken(userId int32) (string, error) {
	return generateJWT(userId, refreshTokenExpireDuration)
}

func GenerateAccessToken(userId int32) (string, error) {
	return generateJWT(userId, accessTokenExpireDuration)
}

func generateJWT(userId int32, exp time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenStr string) (jwt.MapClaims, int) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	switch {
	case token.Valid:
		return token.Claims.(jwt.MapClaims), TokenValid
	case errors.Is(err, jwt.ErrTokenExpired), errors.Is(err, jwt.ErrTokenNotValidYet):
		return nil, TokenExpired
	case errors.Is(err, jwt.ErrTokenMalformed), errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, TokenInvalid
	default:
		return nil, TokenInvalid
	}

}

func saveRefreshToken(userId int32, refreshToken string) error {
	return redisClient.Set(ctx, GetRefreshTokenKey(userId), refreshToken, time.Hour*24*7).Err()
}

func refreshAccessToken(refreshToken string) (string, bool) {
	// 解析refreshToken
	claims, status := ParseJWT(refreshToken)
	if status != TokenValid {
		return "", false
	}

	userId := claims["userId"].(int32)
	newAccessToken, err := generateJWT(userId, accessTokenExpireDuration)
	if err != nil {
		return "", false
	}
	err = redisClient.Set(ctx, GetAccessTokenKey(userId), newAccessToken, accessTokenExpireDuration).Err()
	return newAccessToken, true
}

//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	userId := "exampleUser" // 这应该是经过验证的用户ID
//	accessToken, err := GenerateJWT(userId)
//	if err != nil {
//		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
//		return
//	}
//	refreshToken := uuid.New().String()
//	_ = saveRefreshToken(userId, refreshToken)
//	err = redisClient.Set(ctx, "access:"+userId, accessToken, time.Minute*15).Err()
//	if err != nil {
//		http.Error(w, "Failed to save access token", http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Fprintf(w, "AccessToken: %s\nRefreshToken: %s\n", accessToken, refreshToken)
//}

func validateHandler(w http.ResponseWriter, r *http.Request) bool {
	token := r.FormValue("accessToken")
	if validateAccessToken(token) {
		return true
	} else {
		return false
	}
}

func validateAccessToken(token string) bool {
	_, status := ParseJWT(token)
	return status == TokenValid
}

//func refreshHandler(w http.ResponseWriter, r *http.Request) {
//	refreshToken := r.FormValue("refreshToken")
//	expiredAccessToken := r.FormValue("expiredAccessToken")
//	newAccessToken, err := refreshAccessToken(refreshToken)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusUnauthorized)
//		return
//	}
//	fmt.Fprintf(w, "New AccessToken: %s\n", newAccessToken)
//}
