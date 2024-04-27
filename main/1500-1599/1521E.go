package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1521E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, tot, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &tot, &k)
		a := make([]struct{ c, v int }, k+1)
		for i := 1; i <= k; i++ {
			Fscan(in, &a[i].c)
			a[i].v = i
		}
		sort.Slice(a, func(i, j int) bool { return a[i].c < a[j].c })
		a[0].c = 1e9
		a[0].v = 0
		get := func() int {
			for a[k].c == 0 {
				k--
			}
			a[k].c--
			return a[k].v
		}

		n := 1
		for a[k].c > (n+1)/2*n || tot > n*n-(n/2)*(n/2) {
			n++
		}

		ans := make([][]int, n)
		for i := range ans {
			ans[i] = make([]int, n)
		}
		for i := 0; i < n; i += 2 {
			for j := 1; j < n; j += 2 {
				ans[i][j] = get()
			}
		}
		for i := 0; i < n; i += 2 {
			for j := 0; j < n; j += 2 {
				ans[i][j] = get()
			}
		}
		for i := 1; i < n; i += 2 {
			for j := 0; j < n; j += 2 {
				ans[i][j] = get()
			}
		}

		Fprintln(out, n)
		for _, r := range ans {
			for _, v := range r {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1521E(os.Stdin, os.Stdout) }
