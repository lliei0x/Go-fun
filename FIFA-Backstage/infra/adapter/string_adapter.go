package adapter

import "strings"

// StringClear 格式化字符串 去掉换行
func StringClear(str string) string {
	newStr := strings.TrimSpace(str)
	newReplacer := strings.NewReplacer("\n", "")
	return newReplacer.Replace(newStr)
}
