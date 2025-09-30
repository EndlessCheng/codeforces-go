package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf296A(in io.Reader, out io.Writer) {
	var n, v, mx int
	Fscan(in, &n)
	cnt := [1001]int{}
	for range n {
		Fscan(in, &v)
		cnt[v]++
		mx = max(mx, cnt[v])
	}
	if mx*2-1 <= n {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { cf296A(os.Stdin, os.Stdout) }
