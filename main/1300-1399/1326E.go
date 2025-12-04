package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type data26 struct{ sum, sufMax int }
type seg26 []data26

func (seg26) merge(l, r data26) data26 {
	return data26{l.sum + r.sum, max(l.sufMax+r.sum, r.sufMax)}
}

func (t seg26) update(o, l, r, i, v int) {
	if l == r {
		t[o].sum += v
		t[o].sufMax += v
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, v)
	} else {
		t.update(o<<1|1, m+1, r, i, v)
	}
	t[o] = t.merge(t[o<<1], t[o<<1|1])
}

func cf1326E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, v int
	Fscan(in, &n)
	pos := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		pos[v] = i
	}

	t := make(seg26, 2<<bits.Len(uint(n-1)))
	i := n + 1
	for range n {
		for ; i > 1 && t[1].sufMax <= 0; {
			i--
			t.update(1, 1, n, pos[i], 1)
		}
		Fprint(out, i, " ")
		Fscan(in, &v)
		t.update(1, 1, n, v, -1)
	}
}

//func main() { cf1326E(bufio.NewReader(os.Stdin), os.Stdout) }
