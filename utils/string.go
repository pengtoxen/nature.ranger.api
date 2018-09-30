package utils

import (
	"bytes"
	"strings"
)

//字串截取
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func Strtrim(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}

//字符串连接
func StrConcat(str []string) string {
	buffer := bytes.Buffer{}
	for _, v := range str {
		buffer.WriteString(v)
	}
	return buffer.String()
}
