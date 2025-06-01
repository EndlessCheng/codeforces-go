package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1183H(in io.Reader, out io.Writer) {
	var n, k, ans int
	var s string
	Fscan(in, &n, &k, &s)

	f := [26][]int{}
	for i := range f {
		f[i] = make([]int, n+1)
	}
	sumF := make([]int, n+1)
	for i, b := range s {
		b -= 'a'
		for sz := i + 1; sz > 0; sz-- {
			old := f[b][sz]
			if sz == 1 {
				f[b][sz] = 1
			} else {
				f[b][sz] = min(sumF[sz-1], k)
			}
			sumF[sz] += f[b][sz] - old
		}
	}
	sumF[0] = 1

	for i := n; i >= 0; i-- {
		s := sumF[i]
		if s >= k {
			ans += k * (n - i)
			Fprint(out, ans)
			return
		}
		ans += s * (n - i)
		k -= s
	}
	Fprint(out, -1)
}

//func main() { cf1183H(os.Stdin, os.Stdout) }
