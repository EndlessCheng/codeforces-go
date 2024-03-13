package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
)

// https://space.bilibili.com/206214
type fenwickDiff69 [][]int

func (t fenwickDiff69) add(x, y, val int) {
	for i := x; i < len(t); i += i & -i {
		for j := y; j < len(t[i]); j += j & -j {
			t[i][j] += val
		}
	}
}

func (t fenwickDiff69) update(x1, y1, x2, y2, val int) {
	t.add(x1, y1, val)
	t.add(x1, y2+1, -val)
	t.add(x2+1, y1, -val)
	t.add(x2+1, y2+1, val)
}

func (t fenwickDiff69) get(x, y int) (res int) {
	for i := x; i > 0; i &= i - 1 {
		for j := y; j > 0; j &= j - 1 {
			res += t[i][j]
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
	t := make(fenwickDiff69, n+1)
	for i := range t {
		t[i] = make([]int, m+1)
	}
	for ; q > 0; q-- {
		Fscan(in, &op, &r1, &c1, &r2, &c2)
		if op < 3 {
			t.update(r1, c1, r2, c2, (r1+c1*M+r2*M*M+c2*M*M*M)*(3-op*2))
		} else if t.get(r1, c1) == t.get(r2, c2) {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { cf869E(os.Stdin, os.Stdout) }
