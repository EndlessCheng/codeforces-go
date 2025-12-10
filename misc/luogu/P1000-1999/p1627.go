package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1627(in io.Reader, out io.Writer) {
	var n, b, v, s, ans int
	Fscan(in, &n, &b)
	cnt := map[int]int{0: 1}
	findB := false
	for range n {
		Fscan(in, &v)
		if v == b {
			findB = true
		} else if v < b {
			s--
		} else {
			s++
		}
		if !findB {
			cnt[s]++
		} else {
			ans += cnt[s]
		}
	}
	Fprint(out, ans)
}

//func main() { p1627(bufio.NewReader(os.Stdin), os.Stdout) }
