package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF615D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 1e9 + 7
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

	var k, p int
	e := [2e5]int64{}
	for Fscan(in, &k); k > 0; k-- {
		Fscan(in, &p)
		e[p]++
	}
	sq := true
	for _, c := range e {
		if c%2 > 0 {
			sq = false
			break
		}
	}
	var n, m int64 = 1, 1
	if sq {
		for i, c := range e {
			m = m * (c + 1) % (mod - 1)
			e[i] /= 2
		}
	} else {
		ok := true
		for _, c := range e {
			c++
			if ok && c%2 == 0 {
				c /= 2
				ok = false
			}
			m = m * c % (mod - 1)
		}
	}
	for p, c := range e {
		n = n * pow(int64(p), c*m%(mod-1)) % mod
	}
	Fprint(out, n)
}

//func main() { CF615D(os.Stdin, os.Stdout) }
