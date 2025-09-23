package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1304F2(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	s := make([]int, m+1)
	f := make([]int, m-k+1)
	for i := range n {
		for j := 1; j <= m; j++ {
			Fscan(in, &s[j])
			s[j] += s[j-1]
		}

		if i == 0 {
			for j := range f {
				f[j] = s[j+k] - s[j]
			}
			continue
		}

		nf := make([]int, len(f))
		mx := 0
		q := []int{}
		for j, fj := range f {
			if j >= k {
				mx = max(mx, f[j-k]+s[j]-s[j-k])
			}

			for len(q) > 0 && f[q[len(q)-1]]-s[q[len(q)-1]] <= fj-s[j] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
			if q[0] <= j-k {
				q = q[1:]
			}

			nf[j] = max(mx-s[j], f[q[0]]-s[q[0]]) + s[j+k]
		}

		mx = 0
		q = q[:0]
		for j := m - k; j >= 0; j-- {
			if j <= m-k*2 {
				mx = max(mx, f[j+k]+s[j+k*2]-s[j+k])
			}

			for len(q) > 0 && f[q[len(q)-1]]+s[q[len(q)-1]+k] <= f[j]+s[j+k] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
			if q[0] >= j+k {
				q = q[1:]
			}

			nf[j] = max(nf[j], max(mx+s[j+k], f[q[0]]+s[q[0]+k])-s[j])
		}

		f = nf
	}
	Fprint(out, slices.Max(f))
}

//func main() { cf1304F2(bufio.NewReader(os.Stdin), os.Stdout) }
