package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1066C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var q, i, l int
	r, op := 1, ""
	pos := [2e5 + 1]int{}
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op, &i); op[0] == 'L' {
			pos[i] = l
			l--
		} else if op[0] == 'R' {
			pos[i] = r
			r++
		} else {
			p := pos[i]
			Fprintln(out, min(r-p, p-l)-1)
		}
	}
}

//func main() { CF1066C(os.Stdin, os.Stdout) }
