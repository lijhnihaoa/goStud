/*
	非平凡算法
*/

package example

import "fmt"

func StudGcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	fmt.Println("x: ", x)
	return x
}
