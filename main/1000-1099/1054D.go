package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1054D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	c2 := func(n int) int64 { return int64(n) * int64(n-1) / 2 }

	var n, k, v, s int
	Fscan(in, &n, &k)
	ans := c2(n + 1)
	m := 1<<k - 1
	cnt := map[int]int{s: 1}
	for ; n > 0; n-- {
		Fscan(in, &v)
		s ^= v
		if s > m>>1 {
			s ^= m
		}
		cnt[s]++
	}
	for _, c := range cnt {
		ans -= c2(c/2) + c2(c-c/2)
	}
	Fprint(out, ans)
}

//func main() { CF1054D(os.Stdin, os.Stdout) }
