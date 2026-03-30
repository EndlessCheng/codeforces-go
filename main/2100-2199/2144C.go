package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2144C(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f0, f1 := 1, 0
		preV, preW := 0, 0
		for _, v := range a {
			Fscan(in, &w)
			nf0, nf1 := 0, 0
			if v >= preV && w >= preW {
				nf0 = f0
				nf1 = f1
			}
			if v >= preW && w >= preV {
				nf0 = (nf0 + f1) % mod
				nf1 = (nf1 + f0) % mod
			}
			f0, f1 = nf0, nf1
			preV, preW = v, w
		}
		Fprintln(out, (f0+f1)%mod)
	}
}

//func main() { cf2144C(bufio.NewReader(os.Stdin), os.Stdout) }
