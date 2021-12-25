package example

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv { // 第二个map中的没有第一个mao的key 简单的通过需xv ！= yv判断会有问题
			return false
		}
	}
	return true
}

func StudEqual() {
	flag := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println(flag)
}
