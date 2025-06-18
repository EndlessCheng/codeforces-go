package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2101C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		fa := make([]int, n+1)
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
		f := func() (pos []int) {
			for i := range fa {
				fa[i] = i
			}
			for i, v := range a {
				l := find(v)
				if l > 0 {
					pos = append(pos, i)
					fa[l] = l - 1
				}
			}
			return
		}
		pre := f()
		slices.Reverse(a)
		suf := f()

		ans := 0
		for i, p := range pre {
			ans += max(n-1-suf[i]-p, 0)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2101C(bufio.NewReader(os.Stdin), os.Stdout) }
