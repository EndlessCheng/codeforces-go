package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1700E(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int32, n)
	for i := range a {
		a[i] = make([]int32, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}

	// 好格子：自己是 1，或者存在一个小于自己的邻居
	ok := func(i, j int) bool {
		return a[i][j] == 1 ||
			j > 0 && a[i][j-1] < a[i][j] ||
			j+1 < m && a[i][j+1] < a[i][j] ||
			i > 0 && a[i-1][j] < a[i][j] ||
			i+1 < n && a[i+1][j] < a[i][j]
	}
	// 判断 (i,j)，以及 (i,j) 的邻居，是否都是好格子
	ok2 := func(i, j int) bool {
		return ok(i, j) &&
			(j == 0 || ok(i, j-1)) &&
			(j+1 == m || ok(i, j+1)) &&
			(i == 0 || ok(i-1, j)) &&
			(i+1 == n || ok(i+1, j))
	}

	// 收集坏格子
	type pair struct{ i, j int }
	badPos := []pair{}
	for i := range n {
		for j := range m {
			if !ok(i, j) {
				badPos = append(badPos, pair{i, j})
			}
		}
	}
	if len(badPos) == 0 {
		Fprint(out, 0)
		return
	}

	ans := map[pair]struct{}{}
	// 除了交换 (bi,bj)，也可以通过交换 (bi,bj) 的邻居，使 (bi,bj) 变成一个好格子
	// 只需检查至多 5 个位置，因为 (bi,bj) 必须变成好格子
	bi, bj := badPos[0].i, badPos[0].j
	for _, p := range []pair{{bi, bj}, {bi, bj - 1}, {bi, bj + 1}, {bi - 1, bj}, {bi + 1, bj}} {
		if p.i < 0 || p.i == n || p.j < 0 || p.j == m {
			continue
		}
		for i := range n {
			for j := range m {
				// 交换其他所有点
				a[p.i][p.j], a[i][j] = a[i][j], a[p.i][p.j]
				// 交换离坏格子很远的点，必然是无效交换，所以先检查是否有坏格子仍然是坏格子
				for _, q := range badPos {
					if !ok(q.i, q.j) {
						goto o
					}
				}
				// 有效交换！进一步检查受到影响的 10 个点是否正常
				if ok2(p.i, p.j) && ok2(i, j) {
					// 注意去重
					ans[pair{min(p.i*m+p.j, i*m+j), max(p.i*m+p.j, i*m+j)}] = struct{}{}
				}
			o:
				a[p.i][p.j], a[i][j] = a[i][j], a[p.i][p.j]
			}
		}
	}

	if len(ans) > 0 {
		Fprintln(out, 1, len(ans))
	} else {
		Fprint(out, 2)
	}
}

//func main() { cf1700E(bufio.NewReader(os.Stdin), os.Stdout) }
