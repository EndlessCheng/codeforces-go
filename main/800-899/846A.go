package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF846A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, v, f, c1 int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v == 1 {
			c1++
		} else {
			f = min(f+1, c1)
		}
	}
	Fprint(out, n-f)
}

//func main() { CF846A(os.Stdin, os.Stdout) }
