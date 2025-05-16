package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1778D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		d := 0
		for i := range n {
			if s[i] != t[i] {
				d++
			}
		}
		if d == 0 {
			Fprintln(out, 0)
			continue
		}
		inv := pow(n, mod-2)
		f0, f1 := 0, pow(2, n)-1
		for i := 1; i < d; i++ {
			p := i * inv % mod
			f0, f1 = f1, (f1-1-f0*p%mod)*pow(1-p, mod-2)%mod
		}
		Fprintln(out, (f1+mod)%mod)
	}
}

//func main() { cf1778D(bufio.NewReader(os.Stdin), os.Stdout) }
