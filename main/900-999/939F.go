package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf939F(in io.Reader, out io.Writer) {
	var n, k, l, r int
	Fscan(in, &n, &k)
	f := make([]int, n*2+1)
	for i := 1; i <= n; i++ {
		f[i] = 1e9
	}
	nf := make([]int, n*2+1)
	for range k {
		Fscan(in, &l, &r)
		copy(nf, f)
		q := []int{}
		for j := r; j >= 0; j-- {
			for len(q) > 0 && f[r-j] <= f[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, r-j)
			if q[0] < l-j {
				q = q[1:]
			}
			nf[j] = min(nf[j], f[q[0]]+1)
		}

		q = q[:0]
		for j := range r + 1 {
			for len(q) > 0 && f[j] <= f[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, j)
			if q[0] < j-r+l {
				q = q[1:]
			}
			nf[j] = min(nf[j], f[q[0]]+2)
		}

		f, nf = nf, f
	}

	if f[n] < 1e9 {
		Fprintln(out, "Full")
		Fprint(out, f[n])
	} else {
		Fprint(out, "Hungry")
	}
}

//func main() { cf939F(bufio.NewReader(os.Stdin), os.Stdout) }
