package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1379C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ a, b int }

	var t, N, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &N, &m)
		f := make([]pair, m)
		a := make([]int, m)
		for i := range f {
			Fscan(in, &f[i].a, &f[i].b)
			a[i] = f[i].a
		}
		sort.Slice(f, func(i, j int) bool { a, b := f[i], f[j]; return a.b > b.b || a.b == b.b && a.a > b.a })
		sort.Ints(a)
		ans := int64(0)
		i, sa := m-1, int64(0)
		for _, p := range f {
			for ; N > 0 && i >= 0 && a[i] >= p.b; i-- {
				sa += int64(a[i])
				N--
			}
			n, s := N, sa
			if n > 0 {
				if p.a < p.b {
					s += int64(p.a)
					n--
				}
				s += int64(n) * int64(p.b)
			}
			if s > ans {
				ans = s
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1379C(os.Stdin, os.Stdout) }
