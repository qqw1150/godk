package str

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
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

//
//  Md5 计算字符串的 MD5 散列值
//  @Description: 使用 » RSA 数据安全公司的 MD5 消息摘要算法 计算 string 的 MD5 散列值，并返回该散列值。
//  @param str 要计算的字符串。
//  @return string 以 32 字符的十六进制数形式返回散列值。
//
func Md5(str string) string {
	hash := md5.New()
	io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//
//  Md5File 计算指定文件的 MD5 散列值
//  @Description: 使用 » RSA 数据安全公司的 MD5 消息摘要算法 计算 filename 参数指定的文件的 MD5 散列值，并返回该散列值。该散列值是 32 字符的十六进制数。
//  @param filename 文件名
//  @return string 成功返回字符串，否则返回 空字符串。
//
func Md5File(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	sum := md5.Sum(file)
	return fmt.Sprintf("%x", sum)
}

//
//  Nl2br 在字符串所有新行之前插入 HTML 换行标记
//  @Description: 在字符串 string 所有新行之前插入 '<br />'，并返回。
//  @param str 输入字符串。
//  @return string 返回调整后的字符串。
//
func Nl2br(str string) string {
	p, _ := regexp.Compile("\r\n|\n")
	all := p.ReplaceAll([]byte(str), []byte("<br/>"))
	return string(all)
}

//
//  NumberFormat 以千位分隔符方式格式化一个数字
//  @Description: 以千位分隔符方式格式化一个数字
//  @param number 你要格式化的数字
//  @param decimals 要保留的小数位数
//  @param decPoint 指定小数点显示的字符
//  @param thousandsSep 指定千位分隔符显示的字符
//  @return string 格式化以后的 number.
//
//	example:
//	s:=str.NumberFormat(67657567123456.456,2,".",",")
//	fmt.Println(s)
//	676,575,671,234,56.45
//
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	s := strconv.FormatFloat(number, 'f', decimals, 64)
	arr := strings.Split(s, ".")
	pre := ""
	for k, v := range strings.Split(arr[0], "") {
		if k != 0 && k%3 == 0 {
			pre += thousandsSep
		}
		pre += v
	}
	return pre + decPoint + arr[1]
}

//
//  ParseStr 将字符串解析成多个变量
//  @Description: 如果 string 是 URL 传递入的查询字符串（query string），则将它解析为变量
//  @param str 输入的字符串。
//  @return map[string]interface{} 解析后的变量 （interface类型：string、[]string）
//
//	example:
//	str.ParseStr("first=value&arr[]=foo+bar&arr[]=baz")
//	map[arr:[foo+bar baz] first:value]
//
func ParseStr(str string) map[string]interface{} {
	res := make(map[string]interface{})
	split := strings.Split(str, "&")
	for _, v := range split {
		ps := strings.Split(v, "=")
		if strings.HasSuffix(ps[0], "[]") {

			kk := Substr(ps[0], 0, len(ps[0])-2)
			if res[kk] == nil {
				res[kk] = make([]string, 0)
			}
			tmp := res[kk].([]string)
			tmp = append(tmp, ps[1])
			res[kk] = tmp
		} else {
			res[ps[0]] = ps[1]
		}
	}

	return res
}

//
//  StrStr 查找字符串的首次出现
//  @Description: 返回 haystack 字符串从 needle 第一次出现的位置开始到 haystack 结尾的字符串。
//  @param haystack 输入字符串。
//  @param needle 查找的字符串
//  @param after 若为 true 返回needle后面的字符串，为false返回needle前面的字符串
//  @return string 返回字符串的一部分或者空字符串
//
//	example:
//	s:=str.StrStr("name@example.com","@",false)
//	fmt.Println(s) //name
//	s:=str.StrStr("name@example.com","@",true)
//	fmt.Println(s) //@example.com
//
func StrStr(haystack, needle string, after bool) string {
	index := strings.Index(haystack, needle)
	if index == -1 {
		return ""
	}

	if after {
		return Substr(haystack, index, len(haystack)-index)
	} else {
		return Substr(haystack, 0, index)
	}
}
