package utils

func Difference(a, b []string) []string {
	// 创建一个 map 用于存储切片 B 中的值
	bMap := make(map[string]bool)
	for _, value := range b {
		bMap[value] = true
	}

	// 创建一个结果切片，用于存储切片 A 中比切片 B 多的值
	result := []string{}

	// 遍历切片 A，将不在切片 B 中的值添加到结果切片中
	for _, value := range a {
		if !bMap[value] {
			result = append(result, value)
		}
	}
	return result
}
