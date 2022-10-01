package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF965D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var l, w int
	Fscan(in, &w, &l)
	a := make([]int, w)
	for i := 1; i < w; i++ {
		Fscan(in, &a[i])
		a[i] += a[i-1]
	}
	ans := int(1e9)
	for i := l; i < w; i++ {
		ans = min(ans, a[i]-a[i-l])
	}
	Fprint(out, ans)
}

//func main() { CF965D(os.Stdin, os.Stdout) }
