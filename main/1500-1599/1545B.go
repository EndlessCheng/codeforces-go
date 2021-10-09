package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1545B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		// 将 11 看成一个棋子，可以在棋盘上随意放置
		// 对于连续奇数个 1，在其余 11 位置确定的情况下，剩下的单独的一个 1 的位置也是可以确定的，因此只需考虑 11 的个数
		x, y := strings.Count(s, "11"), strings.Count(s, "0")
		x += y
		num, den := int64(1), int64(1)
		for i := 0; i < y; i++ {
			num = num * int64(x-i) % mod
			den = den * int64(i+1) % mod
		}
		Fprintln(out, num*pow(den, mod-2)%mod)
	}
}

//func main() { CF1545B(os.Stdin, os.Stdout) }
