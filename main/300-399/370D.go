package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// 官方题解枚举了 9 种左上角可能的位置，这样可以省去一堆判断逻辑！

// github.com/EndlessCheng/codeforces-go
func CF370D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	draw := func(x, y, sz int) {
		for i := x; i <= x+sz; i++ {
			for j := y; j <= y+sz; {
				if a[i][j] == '.' {
					a[i][j] = '+'
				}
				if i == x || i == x+sz {
					j++
				} else {
					j += sz
				}
			}
		}
	}
	transpose := func() {
		n, m := len(a), len(a[0])
		b := make([][]byte, m)
		for i := range b {
			b[i] = make([]byte, n)
			for j, r := range a {
				b[i][j] = r[i]
			}
		}
		a = b
	}

	var f func()
	f = func() {
		n, m := len(a), len(a[0])
		i0, i1 := -1, 0
		j0, j1 := m, 0
		for i, row := range a {
			l, r := bytes.IndexByte(row, 'w'), bytes.LastIndexByte(row, 'w')
			if l < 0 {
				continue
			}
			if i0 < 0 {
				i0 = i
			}
			i1 = i
			j0 = min(j0, l)
			j1 = max(j1, r)
		}
		sz := max(i1-i0, j1-j0)
		if sz == 0 {
			return
		}
		if sz >= min(n, m) {
			a = nil
			return
		}
		if sz == i1-i0 {
			col := make([]bool, m)
			for _, r := range a[i0+1 : i1] {
				for j, b := range r {
					if b == 'w' {
						col[j] = true
					}
				}
			}
			pos := []int{}
			for j, b := range col {
				if b {
					pos = append(pos, j)
				}
			}
			if len(pos) == 0 {
				draw(i0, max(j1-sz, 0), sz)
			} else if len(pos) == 1 {
				if pos[0] <= j0 {
					if pos[0]+sz < m {
						draw(i0, pos[0], sz)
						return
					}
				}
				if pos[0] >= j1 {
					if l := pos[0] - sz; l >= 0 {
						draw(i0, l, sz)
						return
					}
				}
				a = nil
			} else if len(pos) == 2 {
				if pos[0] > j0 || pos[1] < j1 || pos[1]-pos[0] != sz {
					a = nil
					return
				}
				draw(i0, pos[0], sz)
			} else {
				a = nil
			}
		} else {
			transpose()
			f()
			if a != nil {
				transpose()
			}
		}
	}
	f()
	if a == nil {
		Fprint(out, -1)
		return
	}
	for _, r := range a {
		Fprintln(out, string(r))
	}
}

//func main() { CF370D(os.Stdin, os.Stdout) }
