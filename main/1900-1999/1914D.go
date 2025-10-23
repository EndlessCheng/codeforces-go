package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1914D(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ v, i int }
		top := [3][]pair{}
		for i := range top {
			a := make([]pair, n)
			for j := range a {
				Fscan(in, &a[j].v)
				a[j].i = j
			}
			slices.SortFunc(a, func(a, b pair) int { return b.v - a.v })
			top[i] = a[:3]
		}

		ans := 0
		for _, p := range top[0] {
			for _, q := range top[1] {
				for _, r := range top[2] {
					if p.i != q.i && p.i != r.i && q.i != r.i {
						ans = max(ans, p.v+q.v+r.v)
					}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1914D(bufio.NewReader(os.Stdin), os.Stdout) }
