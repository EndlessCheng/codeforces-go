package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick []int

func (t fenwick) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] += v
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func cf2143D2(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		row := make([]fenwick, n+1)
		col := make([]fenwick, n+1)
		for i := range row {
			row[i] = make(fenwick, n+1)
			col[i] = make(fenwick, n+1)
		}
		upd := func(x, y, v int) {
			row[y].update(x, v)
			col[x].update(y, v)
		}

		upd(1, 1, 1)
		for range n {
			Fscan(in, &v)
			for y := 1; y <= v; y++ {
				upd(v, y, row[y].pre(v)%mod)
			}
			for x := v + 1; x <= n; x++ {
				upd(x, v, col[x].pre(v)%mod)
			}
		}

		ans := 0
		for _, f := range row {
			ans += f.pre(n)
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { cf2143D2(bufio.NewReader(os.Stdin), os.Stdout) }
