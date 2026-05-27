package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1605E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, x, q int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &x)
		a[i] = x - a[i]
	}
	a[1]++

	b := make([]int, n+1)
	b[1] = 1
	for i := 1; i <= n; i++ {
		for j := i * 2; j <= n; j += i {
			a[j] -= a[i]
			b[j] -= b[i]
		}
	}

	var c, e []int
	s := 0
	for i := 1; i <= n; i++ {
		if b[i] == 0 {
			if a[i] < 0 {
				s += -a[i]
			} else {
				s += a[i]
			}
		} else if b[i] > 0 {
			c = append(c, a[i])
		} else {
			e = append(e, a[i])
		}
	}

	k, t := len(c), len(e)
	slices.Sort(c)
	sc := make([]int, k+1)
	for i, v := range c {
		sc[i+1] = sc[i] + v
	}
	slices.Sort(e)
	se := make([]int, t+1)
	for i, v := range e {
		se[i+1] = se[i] + v
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &x)
		i := sort.Search(k, func(i int) bool { return c[i] > -x })
		j := sort.Search(t, func(i int) bool { return e[i] > x })
		Fprintln(out, s+sc[k]-2*sc[i]+x*(k-2*i)+se[t]-2*se[j]-x*(t-2*j))
	}
}

//func main() { cf1605E(bufio.NewReader(os.Stdin), os.Stdout) }
