package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := 0; i < n; {
		i0 := i
		for i++; i < n && a[i] <= a[i-1]*2; i++ {
		}
		if i-i0 > ans {
			ans = i - i0
		}
	}
	Fprint(out, ans)
}
