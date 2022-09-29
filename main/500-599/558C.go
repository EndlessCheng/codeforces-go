package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF558C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	cnt := [2e5 + 1]int{}
	sum := [1e5 + 1]int{}
	var n, v int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for c := 0; v > 0; v /= 2 {
			cnt[v]++
			sum[v] += c
			c++
		}
	}

	v = 1e5
	for cnt[v] < n {
		v--
	}
	ans := sum[v]
	for v *= 2; cnt[v]*2 > n; v *= 2 {
		ans -= cnt[v]*2 - n
	}
	Fprint(out, ans)
}

//func main() { CF558C(os.Stdin, os.Stdout) }
