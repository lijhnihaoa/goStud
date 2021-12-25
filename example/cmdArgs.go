/*
输出命令行参数
*/
package example

import (
	"fmt"
	"os"
	"strings"
)

// 输出命令行传入的参数
func StudArgsV1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// range 方式读取对值
func StudArgsV2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//使用strings包join进行连接,输出命令行参数
func StudArgsV3() {
	// 输出可执行程序
	// fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
