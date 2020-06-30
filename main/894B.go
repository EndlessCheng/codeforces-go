package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF894B(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var n, m, k int64
	Fscan(in, &n, &m, &k)
	if k == -1 && n&1 != m&1 {
		Fprint(out, 0)
	} else {
		Fprint(out, pow(pow(2, n-1), m-1))
	}
}

//func main() { CF894B(os.Stdin, os.Stdout) }
