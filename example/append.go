package example

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func StudAppend() {
	s := nonempty2([]string{"hello", "", "world"})
	fmt.Println(s)
	fmt.Println([]string{"hello", "", "world"})

	// var x, sliceInt []int

	// for i := 0; i < 10; i++ {
	// 	sliceInt = appendInt(x, i)
	// 	x = sliceInt[:]
	// 	fmt.Printf("%d\tlen=%d\tcap=%d\t%v\n", i, len(x), cap(x), x)
	// }
}
