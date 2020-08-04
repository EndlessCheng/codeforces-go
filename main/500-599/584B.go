package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF584B(_r io.Reader, _w io.Writer) {
	const mod int64 = 1e9 + 7
	pow := func(x int64, n int) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var n int
	Fscan(_r, &n)
	Fprint(_w, (pow(27, n)-pow(7, n)+mod)%mod)
}

//func main() { CF584B(os.Stdin, os.Stdout) }
