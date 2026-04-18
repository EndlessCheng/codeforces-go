package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf981F(in io.Reader, out io.Writer) {
	var n, L int
	Fscan(in, &n, &L)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	b := make([]int, n*3)
	for i := range n {
		Fscan(in, &b[i])
		b[i+n] = b[i] - L
		b[i+n*2] = b[i] + L
	}
	slices.Sort(b)

	ans := sort.Search(L/2, func(x int) bool {
		l, r := 0, n*3-1
		for _, v := range a {
			for b[l] < v-x {
				l++
			}
			for b[r] > v+x {
				r--
			}
			l++
			r++
		}
		return l <= r
	})
	Fprint(out, ans)
}

//func main() { cf981F(bufio.NewReader(os.Stdin), os.Stdout) }
