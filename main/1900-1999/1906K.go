package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1906K(in io.Reader, out io.Writer) {
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

	const w = 17
	var m int
	Fscan(in, &m)
	f := [1 << w]int{}
	for range m {
		var x int
		Fscan(in, &x)
		f[x]++
	}

	for i := range w {
		for j := range 1 << w {
			if j>>i&1 == 0 {
				w0, w1 := f[j], f[j^1<<i]
				f[j] = w0 + w1
				f[j^1<<i] = w0 - w1
			}
		}
	}

	ans := 0
	for _, v := range f {
		ans = (ans + pow(3, (v+m)/2)*pow(mod-1, (m-v)/2)) % mod
	}
	Fprint(out, ans*pow((mod+1)/2, w)%mod)
}

//func main() { cf1906K(bufio.NewReader(os.Stdin), os.Stdout) }
