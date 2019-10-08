package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func Sol5C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s string
	Fscan(in, &s)

	maxLen := map[int]int{0: 1}
	size, pairCnt := 0, 0
	for _, c := range s {
		if c == '(' {
			size++
		} else {
			size--
			if size >= 0 {
				pairCnt++
				maxLen[2*pairCnt]++
			} else {
				size, pairCnt = 0, 0
			}
		}
	}
	ans := 0
	for l := range maxLen {
		if l > ans {
			ans = l
		}
	}
	Fprint(out, ans, maxLen[ans])
}

func main() {
	Sol5C(os.Stdin, os.Stdout)
}
