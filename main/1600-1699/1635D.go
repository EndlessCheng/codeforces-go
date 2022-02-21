package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1635D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7
	var n, p, ans int
	Fscan(in, &n, &p)
	f := make([]int, p+1)
	f[0] = 1
	f[1] = 2
	for i := 2; i <= p; i++ {
		f[i] = (f[i-1] + f[i-2]) % mod
	}

	a := make([]uint, n)
	has := make(map[uint]bool, n)
	for i := range a {
		Fscan(in, &a[i])
		has[a[i]] = true
	}
	if has[1] {
		ans = f[p] - 1
	}
o:
	for _, v := range a {
		for x := v; x&3 != 2; {
			if x == 1 {
				if has[1] {
					continue o
				}
				break
			}
			if x&3 == 0 {
				if x >>= 2; has[x] {
					continue o
				}
			} else {
				if x >>= 1; has[x] {
					continue o
				}
			}
		}
		if q := p - bits.Len(v) + 1; q > 0 {
			ans = (ans + f[q] - 1) % mod
		}
	}
	Fprint(out, (ans+mod)%mod)
}

//func main() { CF1635D(os.Stdin, os.Stdout) }
