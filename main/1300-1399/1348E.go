package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1348E(in io.Reader, out io.Writer) {
	var n, k, a, b, sa, sb int
	Fscan(in, &n, &k)
	f := make([]bool, k)
	f[0] = true
	for range n {
		Fscan(in, &a, &b)
		sa += a
		sb += b
		nf := slices.Clone(f)
		for i := k - 1; i >= 0; i-- {
			if !f[i] {
				continue
			}
			for j := max(k-b, 0); j <= min(a, k); j++ {
				nf[(i+j)%k] = true
			}
		}
		f = nf
	}
	ans := sa/k + sb/k
	sa %= k
	sb %= k
	if k-sb <= sa {
		for _, ok := range f[k-sb : sa+1] {
			if ok {
				ans++
				break
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1348E(os.Stdin, os.Stdout) }
