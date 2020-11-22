package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, m int
	var s, t pair
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	pos := [26][]pair{}
	for i := range a {
		Fscan(in, &a[i])
		for j, b := range a[i] {
			p := pair{i, j}
			if b == 'S' {
				s = p
			} else if b == 'G' {
				t = p
			} else if b != '#' && b != '.' {
				b -= 'a'
				pos[b] = append(pos[b], p)
			}
		}
	}

	type pd struct {
		pair
		d int
	}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	vis[s.x][s.y] = true
	q := []pd{{s, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		if p.pair == t {
			Fprint(out, p.d)
			return
		}
		for _, d := range dir4 {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && a[x][y] != '#' {
				vis[x][y] = true
				q = append(q, pd{pair{x, y}, p.d + 1})
			}
		}
		if b := a[p.x][p.y]; 'a' <= b && b <= 'z' {
			ps := pos[b-'a']
			if ps == nil {
				continue
			}
			for _, tp := range ps {
				if !vis[tp.x][tp.y] {
					vis[tp.x][tp.y] = true
					q = append(q, pd{tp, p.d + 1})
				}
			}
			pos[b-'a'] = nil
		}
	}
	Fprint(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
