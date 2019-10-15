package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1168A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	arr := make([]int, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}
	ans := sort.Search(m, func(op int) bool {
		min := 0
		for _, a := range arr {
			if a+op < m+min {
				if a+op < min {
					return false
				}
				if a > min {
					min = a
				}
			}
		}
		return true
	})
	Fprintln(out, ans)
}

//func main() {
//	Sol1168A(os.Stdin, os.Stdout)
//}
