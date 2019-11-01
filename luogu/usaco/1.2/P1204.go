package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func SolP1204(reader io.Reader, writer io.Writer) {
	type pair struct{ l, r int }
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, ans1, ans2 int
	Fscan(in, &n)
	ps := make([]pair, n)
	for i := range ps {
		Fscan(in, &ps[i].l, &ps[i].r)
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].l < ps[j].l })
	lastR := ps[0].l
	vis := make([]bool, n)
	for i, p := range ps {
		if vis[i] {
			continue
		}
		ans2 = max(ans2, p.l-lastR)
		for j := i + 1; j < n; j++ {
			if vis[j] {
				continue
			}
			if ps[j].l > p.r {
				break
			}
			p.r = max(p.r, ps[j].r)
			vis[j] = true // 忘记写这行 WA 了一发。。
		}
		ans1 = max(ans1, p.r-p.l)
		lastR = p.r
	}
	Fprintln(out, ans1, ans2)
}

//func main() {
//	SolP1204(os.Stdin, os.Stdout)
//}
