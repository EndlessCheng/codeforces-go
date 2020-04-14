package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1335F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := ['Z']pair{
		'U': {-1, 0},
		'D': {1, 0},
		'L': {0, -1},
		'R': {0, 1},
	}
	type tuple struct{ x, y, d int }

	var t, n, m int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		c := make([][]byte, n)
		for i := range c {
			Fscan(in, &c[i])
		}
		inDeg := make([][]int8, n)
		for i := range inDeg {
			inDeg[i] = make([]int8, m)
		}
		rg := make([][][]pair, n)
		for i := range rg {
			rg[i] = make([][]pair, m)
		}
		g := make([][]byte, n)
		for i := range g {
			Fscan(in, &g[i])
			for j, b := range g[i] {
				d := dir4[b]
				xx, yy := i+d.x, j+d.y
				inDeg[xx][yy]++
				rg[xx][yy] = append(rg[xx][yy], pair{i, j})
			}
		}
		visCnt := make([][]int8, n)
		for i := range visCnt {
			visCnt[i] = make([]int8, m)
		}

		f := func(x, y int) {
			stack := []pair{}
			for visCnt[x][y] < 2 {
				// 找基环
				visCnt[x][y]++
				stack = append(stack, pair{x, y})
				d := dir4[g[x][y]]
				x += d.x
				y += d.y
			}
			for i := len(stack) - 1; i >= 0; i-- {
				p := stack[i]
				x, y = p.x, p.y
				if visCnt[x][y] != 1 {
					continue
				}
				// 从树枝上的交叉点开始，反向找树枝上的黑色
				branchBlacks := []int{}
				maxDep := 0
				for q := []tuple{{x, y, 1}}; len(q) > 0; {
					var p tuple
					p, q = q[0], q[1:]
					x, y, d := p.x, p.y, p.d
					visCnt[x][y] = 1 // 访问标记
					inDeg[x][y] = -1 // 访问标记
					if d > maxDep && c[x][y] == '0' {
						branchBlacks = append(branchBlacks, d)
						maxDep = d
					}
					for _, p := range rg[x][y] {
						q = append(q, tuple{p.x, p.y, d + 1})
					}
				}
				// 从环上的交叉点开始，反向找环上的白色
				d := dir4[g[x][y]]
				x += d.x
				y += d.y
				cur := 0
				for dep := 1; dep <= maxDep; dep++ {
					for _, p := range rg[x][y] {
						if visCnt[p.x][p.y] == 2 {
							x, y = p.x, p.y
							break
						}
					}
					if c[x][y] == '1' {
						for ; cur < len(branchBlacks) && branchBlacks[cur] < dep; cur++ {
						}
						if branchBlacks[cur] == dep {
							c[x][y] = '0'
						}
					}
				}
				break
			}
		}
		for i, row := range inDeg {
			for j, d := range row {
				if d == 0 {
					f(i, j)
				}
			}
		}
		for i, row := range visCnt {
			for j, cnt := range row {
				if cnt == 0 {
					f(i, j)
				}
			}
		}

		ans1, ans2 := 0, 0
		for i, row := range visCnt {
			for j, cnt := range row {
				if cnt == 2 {
					ans1++
					if c[i][j] == '0' {
						ans2++
					}
				}
			}
		}
		Fprintln(out, ans1, ans2)
	}
}

//func main() { CF1335F(os.Stdin, os.Stdout) }
