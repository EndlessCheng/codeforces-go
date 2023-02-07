package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1141F2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	type pair struct{ l, r int }
	ans := map[int][]pair{}
	mx := int(1e9)
	for r := 1; r <= n; r++ {
		Fscan(in, &a[r])
		for l, s := r, 0; l > 0; l-- {
			s += a[l]
			ps := ans[s]
			if ps == nil {
				if mx == 1e9 {
					mx = s
				}
				ans[s] = []pair{{l, r}}
			} else if l > ps[len(ps)-1].r {
				if len(ps) >= len(ans[mx]) {
					mx = s
				}
				ans[s] = append(ps, pair{l, r})
			}
		}
	}
	ps := ans[mx]
	Fprintln(out, len(ps))
	for _, p := range ps {
		Fprintln(out, p.l, p.r)
	}
}

//func main() { CF1141F2(os.Stdin, os.Stdout) }
