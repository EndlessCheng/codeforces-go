package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1174C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	ans := make([]int, n+1)
	for i, v := 2, 0; i <= n; i++ {
		if ans[i] == 0 {
			v++
			for j := i; j <= n; j += i {
				ans[j] = v
			}
		}
	}
	for _, v := range ans[2:] {
		Fprint(out, v, " ")
	}
}

//func main() { CF1174C(os.Stdin, os.Stdout) }
