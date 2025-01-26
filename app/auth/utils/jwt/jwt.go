package jwt

import (
	"context"
	"douyin_mall/auth/conf"
	"douyin_mall/auth/utils/redis"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	ctx                        = context.Background()
	jwtSecret                  = []byte(conf.GetConf().Jwt.Secret)
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
	s, err := generateJWT(userId, refreshTokenExpireDuration)
	if err == nil {
		_, err = redis.SetVal(ctx, redis.GetRefreshTokenKey(userId), s, refreshTokenExpireDuration)
		if err != nil {
			return "", err
		}
	}
	return s, err
}

func GenerateAccessToken(userId int32) (string, error) {
	s, err := generateJWT(userId, accessTokenExpireDuration)
	if err == nil {
		_, err = redis.SetVal(ctx, redis.GetAccessTokenKey(userId), s, accessTokenExpireDuration)
		if err != nil {
			return "", err
		}
	}
	return s, err
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

// RefreshAccessToken 刷新access token，同时也要刷新refresh token
func RefreshAccessToken(refreshToken string) (string, string, bool) {
	// 解析refreshToken
	userId, err := GetUserIdFromToken(refreshToken)
	if err != nil {
		klog.Error("userId转换为int失败，", err)
		return "", "", false
	}
	savedRefreshToken, err := redis.GetVal(ctx, redis.GetRefreshTokenKey(userId))
	if err != nil || savedRefreshToken != refreshToken {
		return "", "", false
	}

	newAccessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return "", "", false
	}
	newRefreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", false
	}
	return newAccessToken, newRefreshToken, true
}

func GetUserIdFromToken(token string) (int32, error) {
	claims, status := ParseJWT(token)
	if status != TokenValid {
		return 0, errors.New("token无效")
	}
	userId := int32(claims["userId"].(float64))
	return userId, nil
}
