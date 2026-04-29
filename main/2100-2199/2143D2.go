package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick43 []int

func (t fenwick43) update(i, v int) {
	for ; i < len(t); i += i & -i {
		t[i] += v
	}
}

func (t fenwick43) pre(i int) (res int) {
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
		row := make([]fenwick43, n+1)
		col := make([]fenwick43, n+1)
		for i := range col {
			row[i] = make(fenwick43, n+1)
			col[i] = make(fenwick43, n+1)
		}
		upd := func(x, y, v int) {
			row[x].update(y, v)
			col[y].update(x, v)
		}

		// 定义 f[x][y] 表示表示第一个递增子序列以 x 结尾、第二个递增子序列以 y 结尾时，好子序列的数目
		// 为避免重复统计，规定 x >= y
		upd(1, 1, 1)
		for range n {
			Fscan(in, &v)
			// 把 v 添加到第一个子序列的后面，必须满足 x <= v 且 y <= v
			// 枚举 y，f[v][y] += sum_{x <= v} f[x][y]
			for y := 1; y <= v; y++ {
				upd(v, y, col[y].pre(v)%mod)
			}
			// 把 v 添加到第二个子序列的后面，必须满足 y <= v < x，这里 v != x 从而避免重复统计
			// 枚举 x，f[x][v] += sum_{y <= v} f[x][y]
			for x := v + 1; x <= n; x++ {
				upd(x, v, row[x].pre(v)%mod)
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
