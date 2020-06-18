package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF678D(in io.Reader, out io.Writer) {
	const p int64 = 1e9 + 7
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % p
			}
			x = x * x % p
		}
		return res
	}

	var a, b, n, x int64
	Fscan(in, &a, &b, &n, &x)
	if an := pow(a, n); a == 1 {
		Fprint(out, (an*x+n%p*b)%p)
	} else {
		Fprint(out, (an*x+(an-1)%p*pow(a-1, p-2)%p*b)%p)
	}
}

//func main() { CF678D(os.Stdin, os.Stdout) }
