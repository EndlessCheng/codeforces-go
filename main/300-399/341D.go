package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick41 [][][2][2]int

func (t fenwick41) add(x, y, val int) {
	n := len(t)
	for i := x; i < n; i += i & -i {
		for j := y; j < n; j += j & -j {
			t[i][j][x&1][y&1] ^= val
		}
	}
}

func (t fenwick41) update(x1, y1, x2, y2, val int) {
	t.add(x1, y1, val)
	t.add(x1, y2+1, val)
	t.add(x2+1, y1, val)
	t.add(x2+1, y2+1, val)
}

func (t fenwick41) get(x, y int) (res int) {
	for i := x; i > 0; i &= i - 1 {
		for j := y; j > 0; j &= j - 1 {
			res ^= t[i][j][x&1][y&1]
		}
	}
	return
}

func cf341D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, x1, y1, x2, y2, v int
	Fscan(in, &n, &m)
	t := make(fenwick41, n+1)
	for i := range t {
		t[i] = make([][2][2]int, n+1)
	}
	for range m {
		Fscan(in, &op, &x1, &y1, &x2, &y2)
		if op == 1 {
			Fprintln(out, t.get(x2, y2)^t.get(x2, y1-1)^t.get(x1-1, y2)^t.get(x1-1, y1-1))
		} else {
			Fscan(in, &v)
			t.update(x1, y1, x2, y2, v)
		}
	}
}

//func main() { cf341D(bufio.NewReader(os.Stdin), os.Stdout) }
