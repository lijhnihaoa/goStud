package example

import (
	"bufio"
	"fmt"
	"os"
)

func StudDedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}
}

// input := bufio.NewScanner(os.Stdin)
// input.Split(bufio.ScanWords) //按照单词循环
// input.Scan()
