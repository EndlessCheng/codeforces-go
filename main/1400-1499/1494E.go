package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1494E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	m := map[[2]int]byte{}
	var q, v, w, same, bi int
	var s string
	for Fscan(in, &q, &q); q > 0; q-- {
		if Fscan(in, &s, &v); s[0] == '+' {
			Fscan(in, &w, &s)
			if y := m[[2]int{w, v}]; y > 0 {
				if y == s[0] {
					same++
				}
				bi++
			}
			m[[2]int{v, w}] = s[0]
		} else if s[0] == '-' {
			Fscan(in, &w)
			x := m[[2]int{v, w}]
			delete(m, [2]int{v, w})
			if y := m[[2]int{w, v}]; y > 0 {
				if y == x {
					same--
				}
				bi--
			}
		} else if same > 0 || v&1 > 0 && bi > 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1494E(os.Stdin, os.Stdout) }
