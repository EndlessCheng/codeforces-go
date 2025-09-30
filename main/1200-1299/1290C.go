package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func cf1290C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, m, v, ans int
	var s string
	Fscan(in, &n, &k, &s)
	idx := make([][]int, n)
	fa := make([]int, k+1)
	sz := make([][2]int, k+1)
	sz[0][0] = 1e9 // 强制 x_0=1 的成本为无穷大
	for i := 1; i <= k; i++ {
		Fscan(in, &m)
		for range m {
			Fscan(in, &v)
			idx[v-1] = append(idx[v-1], i)
		}
		fa[i] = i
		// sz[i][0/1] 表示集合中的与 i 开关状态相同/不同的元素个数
		sz[i][0] = 1
	}

	dis := make([]byte, k+1)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			dis[x] ^= dis[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}
	minOp := func(x int) int { return min(sz[x][0], sz[x][1]) }

	for i, p := range idx {
		if p == nil {
			Fprintln(out, ans)
			continue
		}
		from, to := p[0], 0
		if len(p) > 1 {
			to = p[1]
		}
		x, y := find(from), find(to)
		if x != y {
			d := s[i]&1 ^ 1 ^ dis[from] ^ dis[to]
			dis[x] = d
			fa[x] = y
			ans -= minOp(x) + minOp(y)
			sz[y][0] += sz[x][d]
			sz[y][1] += sz[x][d^1]
			ans += minOp(y)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1290C(bufio.NewReader(os.Stdin), os.Stdout) }
