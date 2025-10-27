package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf631C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, tp, r int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	type pair struct{ tp, r int }
	st := []pair{}
	for range m {
		Fscan(in, &tp, &r)
		for len(st) > 0 && st[len(st)-1].r <= r {
			st = st[:len(st)-1]
		}
		if len(st) == 0 || st[len(st)-1].tp != tp {
			st = append(st, pair{tp, r})
		}
	}
	st = append(st, pair{})

	b := slices.Clone(a[:st[0].r])
	slices.Sort(b)
	for i, p := range st[:len(st)-1] {
		l, r := st[i+1].r, p.r
		k := r - l
		if p.tp == 1 {
			copy(a[l:r], b[len(b)-k:])
			b = b[:len(b)-k]
		} else {
			slices.Reverse(b[:k])
			copy(a[l:r], b[:k])
			b = b[k:]
		}
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { cf631C(bufio.NewReader(os.Stdin), os.Stdout) }
