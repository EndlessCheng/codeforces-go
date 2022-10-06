package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF862E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, l, r int
	var v, sa, sb int64
	Fscan(in, &n, &m, &q)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if i&1 > 0 {
			sa -= v
		} else {
			sa += v
		}
	}
	b := make([]int64, m)
	for i := range b[:n] {
		Fscan(in, &b[i])
		if i&1 > 0 {
			sb -= b[i]
		} else {
			sb += b[i]
		}
	}
	f := make([]int64, m-n+1)
	f[0] = sb
	for i := n; i < m; i++ {
		sb = b[i-n] - sb
		Fscan(in, &b[i])
		if n&1 > 0 {
			sb += b[i]
		} else {
			sb -= b[i]
		}
		f[i-n+1] = sb
	}
	sort.Slice(f, func(i, j int) bool { return f[i] < f[j] })

	query := func() {
		i := sort.Search(len(f), func(i int) bool { return f[i] >= sa })
		if i == len(f) || i > 0 && sa-f[i-1] < f[i]-sa {
			Fprintln(out, sa-f[i-1])
		} else {
			Fprintln(out, f[i]-sa)
		}
	}
	query()
	for ; q > 0; q-- {
		Fscan(in, &l, &r, &v)
		if l&1 > 0 {
			if r&1 > 0 {
				sa += v
			}
		} else if r&1 == 0 {
			sa -= v
		}
		query()
	}
}

//func main() { CF862E(os.Stdin, os.Stdout) }
