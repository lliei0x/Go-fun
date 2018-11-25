package adapter

import "strings"

// StringClear 格式化字符串 去掉换行&空格
func StringClear(str string) string {
	// newStr := strings.TrimSpace(str)
	// newReplacer := strings.NewReplacer("\n", "")
	// return newReplacer.Replace(newStr)

	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)

	return str
}
