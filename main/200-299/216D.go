package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF216D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, ans int
	Fscan(in, &n)
	p := make([][]int, n)
	for i := range p {
		Fscan(in, &m)
		p[i] = make([]int, m)
		for j := range p[i] {
			Fscan(in, &p[i][j])
		}
		sort.Ints(p[i])
	}
	for id, b := range p {
		a := p[(id-1+n)%n]
		c := p[(id+1)%n]
		i, k := 0, 0
		for j, p := range b {
			i0 := i
			for i < len(a) && a[i] < p {
				i++
			}
			k0 := k
			for k < len(c) && c[k] < p {
				k++
			}
			if j > 0 && i-i0 != k-k0 {
				ans++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF216D(os.Stdin, os.Stdout) }
