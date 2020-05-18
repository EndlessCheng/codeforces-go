package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	f := func(n, k int, do func([]int)) {
		a := make([]int, k)
		do(a)
		for {
			i := k - 1
			for ; i >= 0; i-- {
				if a[i] != n-1 {
					break
				}
			}
			if i == -1 {
				return
			}
			a[i]++
			for j := i + 1; j < k; j++ {
				a[j] = a[i]
			}
			do(a)
		}
	}

	var n, m, q, ans int
	Fscan(in, &n, &m, &q)
	qs := make([][4]int, q)
	for i := range qs {
		for j := range qs[i] {
			Fscan(in, &qs[i][j])
		}
	}
	f(m, n, func(a []int) {
		c := 0
		for _, q := range qs {
			if a[q[1]-1]-a[q[0]-1] == q[2] {
				c += q[3]
			}
		}
		if c > ans {
			ans = c
		}
	})
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
