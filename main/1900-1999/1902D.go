package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

func cf1902D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dirs := []pair{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}
	var n, q, l, r int
	var s string
	Fscan(in, &n, &q, &s)
	pre := make([]pair, n+1)
	pos := map[pair][]int{{}: {0}}
	p := pair{}
	for i, b := range s {
		d := dirs[b]
		p.x += d.x
		p.y += d.y
		pre[i+1] = p
		pos[p] = append(pos[p], i+1)
	}

	for ; q > 0; q-- {
		Fscan(in, &p.x, &p.y, &l, &r)
		ps := pos[p]
		if len(ps) > 0 && (ps[0] < l || ps[len(ps)-1] >= r) {
			Fprintln(out, "YES")
			continue
		}
		q := pair{pre[r].x + pre[l-1].x - p.x, pre[r].y + pre[l-1].y - p.y}
		ps = pos[q]
		i := sort.SearchInts(ps, l-1)
		if i < len(ps) && ps[i] <= r {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1902D(bufio.NewReader(os.Stdin), os.Stdout) }
