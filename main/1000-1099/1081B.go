package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1081B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, left, color int
	Fscan(in, &n)
	id := make([][]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		id[n-v] = append(id[n-v], i)
	}
	ans := make([]interface{}, n)
	for c := 1; c <= n; c++ {
		if len(id[c])%c > 0 {
			Fprint(out, "Impossible")
			return
		}
		for _, i := range id[c] {
			if left == 0 {
				left = c
				color++
			}
			ans[i] = color
			left--
		}
	}
	Fprintln(out, "Possible")
	Fprintln(out, ans...)
}

//func main() { CF1081B(os.Stdin, os.Stdout) }
