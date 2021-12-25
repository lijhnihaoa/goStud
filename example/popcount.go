package example

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// go 语言中 >> 右移运算比乘除 优先级高
func StudPopCount(x uint64) int {
	fmt.Println("x:", x)
	fmt.Println("pc[byte(x>>0*8)]: ", pc[byte(x>>0*8)], "\tx>>0*8=", x>>0*8)
	fmt.Println("pc[byte(x>>1*8)]: ", pc[byte(x>>1*8)], "\tx>>1*8=", x>>1*8)
	fmt.Println("pc[byte(x>>2*8)]: ", pc[byte(x>>2*8)], "\tx>>2*8=", x>>2*8)
	fmt.Println("pc[byte(x>>3*8)]: ", pc[byte(x>>3*8)], "\tx>>3*8=", x>>3*8)
	fmt.Println("pc[byte(x>>4*8)]: ", pc[byte(x>>4*8)], "\tx>>4*8=", x>>4*8)
	fmt.Println("pc[byte(x>>5*8)]: ", pc[byte(x>>5*8)], "\tx>>5*8=", x>>5*8)
	fmt.Println("pc[byte(x>>6*8)]: ", pc[byte(x>>6*8)], "\tx>>6*8=", x>>6*8)
	fmt.Println("pc[byte(x>>7*8)]: ", pc[byte(x>>7*8)], "\tx>>7*8=", x>>7*8)

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))],
	)
}

func StudPopCountV2(x uint64) int {
	start := time.Now()
	var cnt int
	for i := 0; i < 256; i++ {
		if (byte(x>>i) & 0x1) != 0 {
			cnt++
		}
	}
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
	return cnt
}

func StudPopCountV3(x uint64) int {
	start := time.Now()
	var cnt int
	for {
		if x != 0 {
			cnt++
			x = x & (x - 1)
		} else {
			break
		}
	}
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
	return cnt
}
