package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1351C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := [...]pair{
		'W': {-1, 0},
		'E': {1, 0},
		'S': {0, -1},
		'N': {0, 1},
	}

	var t int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s)
		ans := 0
		vis := map[[2]pair]bool{}
		p := pair{}
		for _, b := range s {
			st := p
			d := dir4[b]
			p.x += d.x
			p.y += d.y
			if pp, qq := [2]pair{st, p}, [2]pair{p, st}; vis[pp] || vis[qq] {
				ans++
			} else {
				vis[pp] = true
				vis[qq] = true
				ans += 5
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1351C(os.Stdin, os.Stdout) }
