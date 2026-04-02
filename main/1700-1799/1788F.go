package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1788F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, v, w, d int
	Fscan(in, &n, &q)
	es := make([]struct{ x, y int }, n-1)
	deg := make([]byte, n)
	for i := range es {
		Fscan(in, &es[i].x, &es[i].y)
		es[i].x--
		es[i].y--
		deg[es[i].x] ^= 1
		deg[es[i].y] ^= 1
	}

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	dis := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			rt := find(fa[x])
			dis[x] ^= dis[fa[x]]
			fa[x] = rt
		}
		return fa[x]
	}
	merge := func(from, to, d int) bool {
		x, y := find(from), find(to)
		if x == y {
			return dis[to]^dis[from] == d
		}
		dis[x] = d ^ dis[to] ^ dis[from]
		fa[x] = y
		return true
	}

	for range q {
		Fscan(in, &v, &w, &d)
		if !merge(v-1, w-1, d) {
			Fprint(out, "No")
			return
		}
	}

	ccOddDegCnt := make([]byte, n)
	xor := 0
	for i, d := range deg {
		find(i)
		if d > 0 { // 每个点权的计算次数是它在原图上的度数，想要影响所有点权的异或值，度数必须是奇数
			ccOddDegCnt[fa[i]] ^= 1 // 连通块中的奇度数点的个数
			xor ^= dis[i]
		}
	}

	for rt, c := range ccOddDegCnt {
		if c > 0 { // 连通块里有奇数个奇度数点，dis[i] ^= xor 会影响答案奇数次，才能真正地影响答案
			for i, f := range fa {
				if f == rt {
					dis[i] ^= xor
				}
			}
			break
		}
	}

	Fprintln(out, "Yes")
	for _, e := range es {
		Fprint(out, dis[e.x]^dis[e.y], " ")
	}
}

//func main() { CF1788F(bufio.NewReader(os.Stdin), os.Stdout) }
