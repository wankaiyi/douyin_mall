package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetString 获取指定环境变量的字符串值
func GetString(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("环境变量 %s 不存在", key)
}

// GetInt 获取指定环境变量的整数值
func GetInt(key string) (int, error) {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return 0, fmt.Errorf("环境变量 %s 的值无法转换为整数: %v", key, err)
		}
		return intValue, nil
	}
	return 0, fmt.Errorf("环境变量 %s 不存在", key)
}

// GetBool 获取指定环境变量的布尔值
func GetBool(key string) (bool, error) {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return false, fmt.Errorf("环境变量 %s 的值无法转换为布尔类型: %v", key, err)
		}
		return boolValue, nil
	}
	return false, fmt.Errorf("环境变量 %s 不存在", key)
}
