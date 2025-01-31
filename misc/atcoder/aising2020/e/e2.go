package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func run2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ k, d int }
		var a, b []pair
		ans := 0
		for i := 0; i < n; i++ {
			Fscan(in, &k, &l, &r)
			if l > r {
				ans += r
				a = append(a, pair{k, l - r})
			} else {
				ans += l
				b = append(b, pair{n - k, r - l})
			}
		}
		f := func(a []pair) {
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
			sort.Slice(a, func(i, j int) bool { return a[i].d > a[j].d })
			for _, p := range a {
				i := find(p.k)
				if i > 0 {
					ans += p.d
					fa[i] = i - 1
				}
			}
		}
		f(a)
		f(b)
		Fprintln(out, ans)
	}
}

//func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
