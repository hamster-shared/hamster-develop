package utils

import (
	"strings"
)

func RemoveDuplicatesAndJoin(input string, split string) string {
	// 将逗号分隔的字符串拆分成切片
	slice := strings.Split(input, split)

	// 创建一个 map 来记录已经出现过的元素
	seen := make(map[string]bool)
	result := []string{}

	// 遍历切片，去除重复项
	for _, item := range slice {
		if item == "" {
			continue
		}
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	// 将切片重新合并为字符串，以逗号分隔
	return strings.Join(result, split)
}

// 判断字符串数组是否包含特定字符串
func ContainsString(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}
