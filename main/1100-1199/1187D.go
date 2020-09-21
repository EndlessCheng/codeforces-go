package main

import (
	"bufio"
	. "fmt"
	"io"
)

func min87(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type node87 struct{ l, r, v int }
type seg87 []node87

func (t seg87) maintain(o int) { t[o].v = min87(t[o<<1].v, t[o<<1|1].v) }
func (t seg87) build(o, l, r int) {
	t[o] = node87{l, r, 1e9}
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}
func (t seg87) update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].v = v
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t.maintain(o)
}
func (t seg87) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].v
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return min87(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

// github.com/EndlessCheng/codeforces-go
func CF1187D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		t := make(seg87, 4*n)
		t.build(1, 1, n)
		pos := make([][]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if pos[v] == nil {
				t.update(1, v, i)
			}
			pos[v] = append(pos[v], i)
		}
		for ; n > 0; n-- {
			Fscan(in, &v)
			if len(pos[v]) == 0 || t.query(1, 1, v) < pos[v][0] {
				for n--; n > 0; n-- {
					Fscan(in, &v)
				}
				Fprintln(out, "NO")
				continue o
			}
			pos[v] = pos[v][1:]
			p := int(1e9)
			if len(pos[v]) > 0 {
				p = pos[v][0]
			}
			t.update(1, v, p)
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1187D(os.Stdin, os.Stdout) }
