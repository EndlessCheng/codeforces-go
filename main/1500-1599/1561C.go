package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1561C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ need, add int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &m)
			mx := 0
			for j := 0; j < m; j++ {
				Fscan(in, &v)
				mx = max(mx, v-j+1)
			}
			a[i] = pair{mx, m}
		}
		sort.Slice(a, func(i, j int) bool { return a[i].need < a[j].need })
		ans, s := 0, 0
		for _, p := range a {
			ans = max(ans, p.need-s)
			s += p.add
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1561C(os.Stdin, os.Stdout) }
