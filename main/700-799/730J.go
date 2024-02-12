package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF730J(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, sa, sb, m, maxSave int
	Fscan(in, &n)
	a := make([]struct{ a, b int }, n)
	for i := range a {
		Fscan(in, &a[i].a)
		sa += a[i].a
	}
	for i := range a {
		Fscan(in, &a[i].b)
	}

	sort.Slice(a, func(i, j int) bool { return a[i].b > a[j].b })
	for i, p := range a {
		sb += p.b
		if sb >= sa {
			m = i + 1
			break
		}
	}

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, sb+1)
		for j := range f[i] {
			f[i][j] = -1e9
		}
	}
	f[0][0] = 0
	s := 0
	for i := len(a) - 1; i >= 0; i-- {
		p := a[i]
		s = min(s+p.b, sb)
		for j := m; j > 0; j-- {
			for k := s; k >= p.b; k-- {
				f[j][k] = max(f[j][k], f[j-1][k-p.b]+p.a)
			}
		}
	}
	for _, save := range f[m][sa:] {
		maxSave = max(maxSave, save)
	}
	Fprint(out, m, sa-maxSave)
}

//func main() { CF730J(os.Stdin, os.Stdout) }
