package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e9
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, m, t, sx, sy, tx, ty, ans int
	Fscan(in, &n, &m, &t)
	a := make([]string, n)
	ps := []pair{}
	for i := range a {
		Fscan(in, &a[i])
		for j, c := range a[i] {
			if c == 'S' {
				sx, sy = i, j
			} else if c == 'G' {
				tx, ty = i, j
			} else if c == 'o' {
				ps = append(ps, pair{i, j})
			}
		}
	}
	ps = append(ps, pair{sx, sy}, pair{tx, ty})
	k := len(ps)
	idx := make(map[pair]int, k)
	for i, p := range ps {
		idx[p] = i
	}

	vis := make([][]int, n)
	for i := range vis {
		vis[i] = make([]int, m)
	}
	dis := make([][]int, k)
	for i, p := range ps {
		dis[i] = make([]int, k)
		for j := range dis[i] {
			dis[i][j] = inf
		}
		q := []pair{p}
		vis[p.x][p.y] = i + 1
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dir4 {
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < n && 0 <= y && y < m && vis[x][y] != i+1 && a[x][y] != '#' {
						if a[x][y] != '.' {
							dis[i][idx[pair{x, y}]] = step
						}
						vis[x][y] = i + 1
						q = append(q, pair{x, y})
					}
				}
			}
		}
	}
	if dis[k-2][k-1] > t {
		Fprint(out, -1)
		return
	}

	k -= 2 // 下面 k 表示起点，k+1 表示终点
	f := make([][]int, 1<<k)
	for i := range f {
		f[i] = make([]int, k)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	for i, d := range dis[k][:k] {
		f[1<<i][i] = d
	}
	for s, dr := range f {
		for _s := uint(s); _s > 0; _s &= _s - 1 {
			i := bits.TrailingZeros(_s)
			if dr[i]+dis[k+1][i] <= t {
				ans = max(ans, bits.OnesCount(uint(s)))
			}
			for cus, lb := len(f)-1^s, 0; cus > 0; cus ^= lb {
				lb = cus & -cus
				ns := s | lb
				j := bits.TrailingZeros(uint(lb))
				f[ns][j] = min(f[ns][j], dr[i]+dis[i][j])
			}
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
func max(a, b int) int { if b > a { return b }; return a }
