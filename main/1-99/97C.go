package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf97C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	p := make([]float64, n+1)
	for i := range p {
		Fscan(in, &p[i])
	}
	m := float64(n) / 2
	ans := 0.
	for i := range n/2 + 1 {
		for j := n/2 + 1; j <= n; j++ {
			ans = max(ans, (p[i]*(float64(j)-m)+p[j]*(m-float64(i)))/float64(j-i))
		}
	}
	Fprintf(out, "%.6f", ans)
}

//func main() { cf97C(os.Stdin, os.Stdout) }
