package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF696C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	const mod2 = (mod + 1) / 2
	const mod3 = (mod + 1) / 3
	pow := func(x, n int64) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n int
	var v int64
	q, even := int64(2), false
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		q = pow(q, v)
		if v&1 == 0 {
			even = true
		}
	}
	// 实际上的幂次是 n-1
	q = q * mod2 % mod
	p := q
	if even {
		p++
	} else {
		p += mod - 1
	}
	Fprint(out, p*mod3%mod, "/", q)
}

//func main() { CF696C(os.Stdin, os.Stdout) }
