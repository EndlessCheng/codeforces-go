package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1661C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		mx := 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] > mx {
				mx = a[i]
			}
		}
		f := func(t int) int64 {
			var c1, c2 int64
			for _, v := range a {
				v = t - v
				c1 += int64(v % 2)
				c2 += int64(v / 2)
			}
			if c1 > c2 {
				return c1*2 - 1
			}
			if c1 == c2 {
				return c1 * 2
			}
			// 会出现 12_2_2_2_2_2 的情况
			// 但是两个空闲的 1 可以组成一个 2
			c2 = (c2 - c1) * 2
			return c1*2 + c2/3*2 + c2%3 // 手玩
		}
		Fprintln(out, min(f(mx), f(mx+1)))
	}
}

//func main() { CF1661C(os.Stdin, os.Stdout) }
