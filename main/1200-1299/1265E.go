package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1265E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	const mod int64 = 998244353
	var exgcd func(a, b int64) (gcd, x, y int64)
	exgcd = func(a, b int64) (gcd, x, y int64) {
		if b == 0 {
			return a, 1, 0
		}
		gcd, y, x = exgcd(b, a%b)
		y -= a / b * x
		return
	}
	modInverse := func(a, m int64) int64 {
		_, x, _ := exgcd(a, m)
		res := (x%m + m) % m
		return res
	}
	modFrac := func(a, b, m int64) int64 {
		return a * modInverse(b, m) % m
	}

	var n int
	Fscan(in, &n)
	var ans, v int64
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		ans = (ans + 1) * modFrac(100, v, mod) % mod
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1265E(os.Stdin, os.Stdout)
//}
