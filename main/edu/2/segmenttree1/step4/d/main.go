package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, i int
	var v int8
	Fscan(in, &n, &q)
	trees := [41][]int{}
	for i := 1; i <= 40; i++ {
		trees[i] = make([]int, n+1)
	}
	a := make([]int8, n+1)
	for p := 1; p <= n; p++ {
		Fscan(in, &a[p])
		t := trees[a[p]]
		for i := p; i <= n; i += i & -i {
			t[i]++
		}
	}
	set := func(p int, val int8) {
		t := trees[a[p]]
		for i := p; i <= n; i += i & -i {
			t[i]--
		}
		t = trees[val]
		for i := p; i <= n; i += i & -i {
			t[i]++
		}
		a[p] = val
	}
	query := func(L, R int) (cnt int8) {
		for i := 1; i <= 40; i++ {
			t := trees[i]
			s, l, r := 0, L-1, R
			for ; r > l; r &= r - 1 {
				s += t[r]
			}
			for ; l > r; l &= l - 1 {
				s -= t[l]
			}
			if s > 0 {
				cnt++
			}
		}
		return
	}
	for ; q > 0; q-- {
		Fscan(in, &op)
		if op == 1 {
			Fscan(in, &l, &r)
			Fprintln(out, query(l, r))
		} else {
			Fscan(in, &i, &v)
			set(i, v)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
