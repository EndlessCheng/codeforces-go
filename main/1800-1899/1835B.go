package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1835B(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	p := []int{}
	add := func(x int) {
		for i := x - 2; i <= x+2; i++ {
			if 0 <= i && i <= m && (len(p) == 0 || i > p[len(p)-1]) {
				p = append(p, i)
			}
		}
	}

	add(0)
	for _, x := range a {
		add(x)
	}
	add(m)

	mx, ans := -1, 0
	i, j := 0, 0
	for _, x := range p {
		for i < n && x > a[i] {
			i++
		}
		for j < n && x >= a[j] {
			j++
		}
		l, r := 0, m
		if j >= k {
			l = (a[j-k]+x)/2 + 1
		}
		if i+k <= n {
			r = (a[i+k-1] + x - 1) / 2
		}
		if r-l+1 > mx {
			mx = r - l + 1
			ans = x
		}
	}
	Fprint(out, mx, ans)
}

//func main() { cf1835B(bufio.NewReader(os.Stdin), os.Stdout) }
