package main

import (
	"bytes"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	const mx = 61
	C := [mx][mx]int{}
	for i := 0; i < mx; i++ {
		C[i][0], C[i][i] = 1, 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	var a, b, k int
	Fscan(in, &a, &b, &k)
	ans := bytes.Repeat([]byte{'a'}, a+b)
	for i := range ans {
		if k > C[a+b-1][b] {
			k -= C[a+b-1][b]
			ans[i] = 'b'
			b--
		} else {
			a--
		}
	}
	Fprintf(out, "%s", ans)
}

func main() { run(os.Stdin, os.Stdout) }
