package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1630C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, curR, nxtR int
	Fscan(in, &n)
	a := make([]int, n)
	r := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		r[a[i]] = i
	}

	ans := n
	for i, v := range a {
		if i > nxtR {
			curR = i
		}
		nxtR = max(nxtR, r[v])
		if i == curR {
			curR = nxtR
			ans--
		}
	}
	Fprint(out, ans)
}

//func main() { CF1630C(os.Stdin, os.Stdout) }
