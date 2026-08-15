package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2254E(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		d := make([]int, n)
		s := 0
		for i := range d {
			Fscan(in, &d[i])
			s += d[i]
		}
		if s < 1 {
			Fprintln(out, -1)
			continue
		}
		slices.Sort(d)

		fa := make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
		find := func(x int) int {
			rt := x
			for fa[rt] != rt {
				rt = fa[rt]
			}
			for fa[x] != rt {
				fa[x], x = rt, fa[x]
			}
			return rt
		}

		s = 0
		ans := []any{}
		for _, v := range d[sort.SearchInts(d, 1):] {
			s += v
			ans = append(ans, s)
			for {
				j := find(sort.SearchInts(d, 1-s))
				if j == n || d[j] > 0 {
					break
				}
				s += d[j]
				ans = append(ans, s)
				fa[j] = find(j + 1)
			}
		}
		Fprintln(out, ans...)
	}
}

//func main() { cf2254E(bufio.NewReader(os.Stdin), os.Stdout) }
