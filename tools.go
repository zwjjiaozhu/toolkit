package toolkit

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// FirstLetter Capitalize
func FirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// CutStr 切割字符串，继续优化
func CutStr(s string, start, end int) string {
	// tod: -索引
	var length = len(s)
	if length == 0 {
		return s
	}
	// 合法区间
	if start < 0 {
		start = 0
	} else if start > length-1 {
		start = length - 1
	}
	if end > length-1 {
		end = length - 1
	}
	return s[start:end]
}

// SliceContains 切片包含
func SliceContains[T int | string](arr []T, s T) bool {
	for i := 0; i < len(arr); i++ {
		if s == arr[i] {
			return true
		}
	}
	return false
}

// StrLen 字符数
// 计算unicode的数量
//func StrLen(text string) int {
//	return len([]rune(text))
//}

// MergeMap 合并字典
func MergeMap[T int | string](a, b map[T]T) map[T]T {
	for k, v := range b {
		a[k] = v
	}
	return a
}

// UpdateMap 用字典b的值更新字典a
func UpdateMap[T int | string](a, b map[T]T) map[T]T {
	for k := range a {
		if v, ok := b[k]; ok {
			a[k] = v
		}
	}
	return a
}

func Max[T int | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T int | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// IfElse 三元运算符
func IfElse[T string](condition bool, a, b T) T {
	if condition {
		return a
	} else {
		return b
	}
}

func MD5(text string) string {
	has := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", has)
}
