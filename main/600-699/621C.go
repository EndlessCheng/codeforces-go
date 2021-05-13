package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF621C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, l, r int
	Fscan(in, &n, &m)
	p := make([]float64, n)
	for i := range p {
		Fscan(in, &l, &r)
		l--
		p[i] = 1 - float64(r/m-l/m)/float64(r-l)
	}
	ans := .0
	for i, v := range p {
		ans += 1 - v*p[(i+1)%n]
	}
	Fprintf(out, "%.8f", ans*2000)
}

//func main() { CF621C(os.Stdin, os.Stdout) }
