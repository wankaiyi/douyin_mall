package jwt

import (
	"context"
	"douyin_mall/auth/conf"
	"douyin_mall/auth/infra/rpc"
	"douyin_mall/auth/utils/redis"
	"douyin_mall/rpc/kitex_gen/user"
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
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(refreshTokenExpireDuration).Unix(),
	}
	s, err := generateJWT(claims)
	if err == nil {
		_, err = redis.SetVal(ctx, redis.GetRefreshTokenKey(userId), s, refreshTokenExpireDuration)
		if err != nil {
			return "", err
		}
	}
	return s, err
}

func GenerateAccessToken(userId int32, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(accessTokenExpireDuration).Unix(),
	}
	s, err := generateJWT(claims)
	if err == nil {
		_, err = redis.SetVal(ctx, redis.GetAccessTokenKey(userId), s, accessTokenExpireDuration)
		if err != nil {
			return "", err
		}
	}
	return s, err
}

func generateJWT(claims jwt.MapClaims) (string, error) {
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

// RefreshAccessToken 刷新access token，同时也要刷新refresh token;；
// 调用场景：1. access token 过期，需要刷新；2. 用户角色变更，后端将access token删除，让前端主动刷新access token
func RefreshAccessToken(refreshToken string) (string, string, bool) {
	// 解析refreshToken
	claims, status := ParseJWT(refreshToken)
	if status != TokenValid {
		klog.Error("refreshToken无效，解析失败，refreshToken: %s", refreshToken)
		return "", "", false
	}
	userId := int32(claims["userId"].(float64))
	getUserRoleByIdResp, err := rpc.UserClient.GetUserRoleById(ctx, &user.GetUserRoleByIdReq{
		UserId: userId,
	})
	if err != nil {
		klog.Errorf("rpc调用GetUserRoleById失败，userId: %d, err: %v", userId, err)
		return "", "", false
	}
	role := getUserRoleByIdResp.Role

	savedRefreshToken, err := redis.GetVal(ctx, redis.GetRefreshTokenKey(userId))
	if err != nil || savedRefreshToken != refreshToken {
		return "", "", false
	}

	newAccessToken, err := GenerateAccessToken(userId, role)
	if err != nil {
		return "", "", false
	}
	newRefreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", false
	}
	return newAccessToken, newRefreshToken, true
}
