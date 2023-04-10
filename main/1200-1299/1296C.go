package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1296C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int }
	dir4 := [...]pair{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		l, r := -1, n
		p := pair{}
		last := map[pair]int{p: 0}
		for i, c := range s {
			d := dir4[c]
			p.x += d.x
			p.y += d.y
			if j, ok := last[p]; ok && i-j < r-l {
				l, r = j, i
			}
			last[p] = i + 1
		}
		if l < 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, l+1, r+1)
		}
	}
}

//func main() { CF1296C(os.Stdin, os.Stdout) }
