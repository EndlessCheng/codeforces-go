package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1693D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	ans := 0
	f := make([]int, n+1)
	g := make([]int, n+1)
	r := n
	for l := n; l > 0; l-- {
		f[l] = n + 1
		for i := l + 1; i <= n; i++ {
			nf, ng := 0, n+1
			if a[i] > a[i-1] {
				nf = max(nf, f[i-1])
			}
			if a[i] > g[i-1] {
				nf = max(nf, a[i-1])
			}
			if a[i] < a[i-1] {
				ng = min(ng, g[i-1])
			}
			if a[i] < f[i-1] {
				ng = min(ng, a[i-1])
			}
			if nf == f[i] && ng == g[i] {
				break
			}
			f[i], g[i] = nf, ng
		}

		for f[r] == 0 && g[r] > n {
			r--
		}
		ans += r - l + 1
	}
	Fprint(out, ans)
}

//func main() { cf1693D(bufio.NewReader(os.Stdin), os.Stdout) }
