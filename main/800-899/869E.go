package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
)

// https://space.bilibili.com/206214
type fenwick69 [][]int

func (f fenwick69) add(x, y, val int) {
	for i := x; i < len(f); i += i & -i {
		for j := y; j < len(f[i]); j += j & -j {
			f[i][j] += val
		}
	}
}

func (f fenwick69) get(x, y int) (res int) {
	for i := x; i > 0; i &= i - 1 {
		for j := y; j > 0; j &= j - 1 {
			res += f[i][j]
		}
	}
	return
}

func cf869E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	M := rand.Intn(1e9) + 1e9

	var n, m, q, op, r1, c1, r2, c2 int
	Fscan(in, &n, &m, &q)
	t := make(fenwick69, n+1)
	for i := range t {
		t[i] = make([]int, m+1)
	}
	for ; q > 0; q-- {
		Fscan(in, &op, &r1, &c1, &r2, &c2)
		if op < 3 {
			v := (r1 + c1*M + r2*M*M + c2*M*M*M) * (3 - op*2)
			t.add(r1, c1, v)
			t.add(r1, c2+1, -v)
			t.add(r2+1, c1, -v)
			t.add(r2+1, c2+1, v)
		} else if t.get(r1, c1) == t.get(r2, c2) {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { cf869E(os.Stdin, os.Stdout) }
