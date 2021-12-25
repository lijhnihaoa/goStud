package example

import (
	"fmt"
	"strings"
)

//仅仅支持整数
func StudComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return StudComma(s[:n-3]) + "," + s[n-3:]
}

//处理小数部分
func commaFloat(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return s[:3] + "," + commaFloat(s[3:])
}

// 支持浮点数
func StudCommaV2(s string) string {
	var stringInt string
	var stringFloat string
	dot := strings.LastIndex(s, ".")
	if dot != -1 {

		stringInt = s[:dot]
		stringFloat = s[dot+1:]
	} else {
		stringInt = s
		return StudComma(stringInt)
	}
	fmt.Println("stringInt: ", stringInt)
	fmt.Println("stringFloat: ", stringFloat)

	stringInt = StudComma(stringInt)
	stringFloat = commaFloat(stringFloat)

	return stringInt + "." + stringFloat
}
