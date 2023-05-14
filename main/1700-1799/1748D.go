package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1748D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var exgcd func(a, b int64) (gcd, x, y int64)
	exgcd = func(a, b int64) (gcd, x, y int64) {
		if b == 0 {
			return a, 1, 0
		}
		gcd, y, x = exgcd(b, a%b)
		y -= a / b * x
		return
	}
	// ax-b 是 m 的倍数
	inv := func(a, b, m int64) int64 {
		g, x, _ := exgcd(a, m)
		if b%g != 0 {
			return -1
		}
		x *= b / g
		m /= g
		return (x%m + m) % m
	}

	var T int
	var a, b, d int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &d)
		c := a | b
		x := inv(1<<30, -c, d)
		if x < 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, x<<30|c)
		}
	}
}

//func main() { CF1748D(os.Stdin, os.Stdout) }
