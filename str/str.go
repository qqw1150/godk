package str

import (
	"strings"
)

//
//  Substr 返回字符串的子串
//  @Description: 返回字符串 str 由 start 和 length 参数指定的子字符串。
//  @param str 输入字符串。必须至少有一个字符。
//  @param start 如果 start 是非负数，返回的字符串将从 string 的 start 位置开始，从 0 开始计算。例如，在字符串 "abcdef" 中，在位置 0 的字符是 "a"，位置 2 的字符串是 "c" 等等。
//如果 start 是负数，返回的字符串将从 string 结尾处向前数第 start 个字符开始。
//如果 string 的长度小于 start，将返回 空字符串。
//  @param length 截取字符串的长度
//  @return string 返回提取的子字符串
func Substr(str string, start int, length int) string {
	strLen := len(str)
	if strLen <= 0 {
		return ""
	}

	if start < 0 {
		if strLen+start >= 0 {
			start = strLen + start
		} else {
			start = 0
		}

	} else if start >= strLen {
		start = strLen - 1
	}

	if length >= strLen-start {
		length = strLen - start
	} else if length < 0 {
		length = 0
	}

	res := ""
	count := start + length
	arr := strings.Split(str, "")
	for ; start < count; start++ {
		res += arr[start]
	}
	return res
}

//
//  Ucfirst 将字符串的首字母转换为大写
//  @Description: 将 str 的首字符（如果首字符是字母）转换为大写字母，并返回这个字符串。
//  @param str 输入字符串。
//  @return string 返回结果字符串。
//
func Ucfirst(str string) string {
	if len(str) <= 0 {
		return ""
	}

	arr := strings.Split(str, "")
	arr[0] = strings.ToUpper(arr[0])
	return strings.Join(arr, "")
}

//
//  Ucwords 将字符串中每个单词的首字母转换为大写
//  @Description: 将 str 中每个单词的首字符（如果首字符是字母）转换为大写字母，并返回这个字符串。
//  @param str 输入字符串。
//  @return string 返回转换后的字符串。
//
func Ucwords(str string) string {
	if len(str) <= 0 {
		return ""
	}

	runes := []rune(str)
	pk := 0
	for k, r := range runes {
		if (k == 0 || strings.ContainsRune("\r\b\f\v\t ", runes[pk])) && r >= 97 && r <= 123 {
			runes[k] = r - 32
		} else {
			runes[k] = r
		}
		pk = k
	}

	return string(runes)
}

//
//  UcwordsForDelimiters 将字符串中每个单词的首字母转换为大写
//  @Description:  将 str 中每个单词的首字符（如果首字符是字母）转换为大写字母，并返回这个字符串。
//  @param str 输入字符串。
//  @param delimiters 单词分割字符。
//  @return string 返回转换后的字符串。
//
func UcwordsForDelimiters(str, delimiters string) string {
	if len(str) <= 0 {
		return ""
	}

	runes := []rune(str)
	pk := 0
	for k, r := range runes {
		if (k == 0 || strings.ContainsRune(delimiters, runes[pk])) && r >= 97 && r <= 123 {
			runes[k] = r - 32
		} else {
			runes[k] = r
		}
		pk = k
	}

	return string(runes)
}
