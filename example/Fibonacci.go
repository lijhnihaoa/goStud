/*
	斐波那契数列
*/

package example

import "fmt"

func StudFib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	fmt.Println("x: ", x)
	return x
}
