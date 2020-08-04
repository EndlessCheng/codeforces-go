package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF690D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e6 + 3
	fact := func(n int) int64 {
		x := int64(1)
		for i := int64(2); i <= int64(n); i++ {
			x = x * i % mod
		}
		return x
	}
	pow := func(x int64, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	inv := func(a int64) int64 { return pow(a, mod-2) }
	frac := func(a, b int64) int64 { return a * inv(b) % mod }

	var n, c int
	Fscan(in, &n, &c)
	Fprint(out, frac(fact(n+c), fact(n)*fact(c)%mod)-1)
}

//func main() { CF690D2(os.Stdin, os.Stdout) }
