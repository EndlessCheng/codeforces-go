package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2069C(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var f1, f2, f3 int
		for range n {
			Fscan(in, &v)
			if v == 1 {
				f1++
			} else if v == 2 {
				f2 = (f2*2 + f1) % mod
			} else {
				f3 += f2
			}
		}
		Fprintln(out, f3%mod)
	}
}

//func main() { cf2069C(bufio.NewReader(os.Stdin), os.Stdout) }
