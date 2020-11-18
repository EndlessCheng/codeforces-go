package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1439A2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]byte, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := [][][]int{}
		add := func(ps [][]int) {
			ans = append(ans, ps)
			for _, p := range ps {
				a[p[0]][p[1]] ^= 1
			}
		}
		for i := n - 1; i > 1; i-- {
			for j, b := range a[i] {
				if b == '1' {
					if j+1 < m {
						add([][]int{{i, j}, {i - 1, j}, {i - 1, j + 1}})
					} else {
						add([][]int{{i, j}, {i - 1, j}, {i - 1, j - 1}})
					}
				}
			}
		}
		for j := m - 1; j > 1; j-- {
			for i := 0; i < 2; i++ {
				if a[i][j] == '1' {
					add([][]int{{i, j}, {0, j - 1}, {1, j - 1}})
				}
			}
		}

		cell := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
		var p0, p1 []int
		for i, p := range cell {
			if a[p[0]][p[1]] == '0' {
				p0 = append(p0, i)
			} else {
				p1 = append(p1, i)
			}
		}
		if len(p1) == 4 {
			ans = append(ans, [][]int{cell[p1[0]], cell[p1[1]], cell[p1[2]]})
			p0, p1 = p1[:3], p1[3:]
		}
		if len(p1) == 1 {
			ans = append(ans, [][]int{cell[p1[0]], cell[p0[0]], cell[p0[1]]})
			p1, p0 = p0[:2], append(p0[2:], p1[0])
		}
		if len(p1) == 2 {
			ans = append(ans, [][]int{cell[p1[0]], cell[p0[0]], cell[p0[1]]})
			p1 = append(p1[1:], p0...)
		}
		if len(p1) == 3 {
			ans = append(ans, [][]int{cell[p1[0]], cell[p1[1]], cell[p1[2]]})
		}

		Fprintln(out, len(ans))
		for _, ps := range ans {
			for _, p := range ps {
				Fprint(out, p[0]+1, p[1]+1, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1439A2(os.Stdin, os.Stdout) }
