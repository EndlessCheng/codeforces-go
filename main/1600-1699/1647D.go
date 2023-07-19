package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1647D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	isPrime := func(n int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	var T, x, d int
	f := func() bool {
		k := 0
		for ; x%d == 0; x /= d {
			k++
		}
		return k > 1 && // 如果 k = 1，那么方案只有一种，x 自己，不符合要求
			   (!isPrime(x) || // x 是合数，有多种分解
				// 此时 x 是 1 或质数
				// 需要拿一个 d 出来，分解给 x * 1 * 1 * ...
				// 那么 d 必须是合数，且 k >= 3
				// 特别地，如果 k = 3，那么 d 不能分解成 x^2（注意 x 是质数），否则 x 分给 x 得到 x^2=d，不符合要求
				!(k == 2 || k == 3 && int64(d) == int64(x)*int64(x) || isPrime(d)))
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &d)
		if f() {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1647D(os.Stdin, os.Stdout) }
