package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1661D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var n, k int
	Fscan(in, &n, &k)
	b := make([]int64, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	var ans, a, d int64
	d2 := make([]int64, n) // 二阶差分
	for i := n - 1; i >= 0; i-- {
		d += d2[i] // 一阶差分
		a += d // a[i] 的值
		if a < b[i] {
			k2 := min(i+1, k)
			times := (b[i]-a-1)/int64(k2) + 1
			ans += times
			a += times * int64(k2)
			if i > 0 {
				d2[i-1] -= times
				if i > k2 {
					d2[i-k2-1] += times
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1661D(os.Stdin, os.Stdout) }
