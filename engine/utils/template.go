package utils

import (
	"regexp"
	"strings"
)

// ReplaceWithParam 参数替换
func ReplaceWithParam(content string, paramMap map[string]string) string {
	compileRegex := regexp.MustCompile("\\${{.*?}}")
	matchArr := compileRegex.FindStringSubmatch(content)

	result := strings.Clone(content)
	for _, arr := range matchArr {
		key := strings.Replace(arr, "${{", "", 1)
		key = strings.Replace(key, "}}", "", 1)
		key = strings.Fields(key)[0]
		if strings.HasPrefix(key, "param.") {
			key = strings.Replace(key, "param.", "", 1)
			result = strings.Replace(result, arr, paramMap[key], -1)
		}
	}

	return result
}
