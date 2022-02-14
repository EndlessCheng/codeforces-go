package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF8B(in io.Reader, out io.Writer) {
	type pair struct{ x, y int }
	dir4 := []pair{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}}
	var s string
	Fscan(in, &s)
	p := map[pair]int{{}: -1}
	x, y := 0, 0
	for i, b := range s {
		d := dir4[b]
		x += d.x
		y += d.y
		if _, ok := p[pair{x, y}]; ok {
			Fprint(out, "BUG")
			return
		}
		for _, d := range dir4 {
			if j, ok := p[pair{x + d.x, y + d.y}]; ok && i-j > 1 {
				Fprint(out, "BUG")
				return
			}
		}
		p[pair{x, y}] = i
	}
	Fprint(out, "OK")
}

//func main() { CF8B(os.Stdin, os.Stdout) }
