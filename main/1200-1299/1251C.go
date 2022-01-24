package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1251C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		a := [2][]byte{}
		for _, b := range s {
			a[b&1] = append(a[b&1], b)
		}
		s = s[:0]
		i, j := 0, 0
		for i < len(a[0]) && j < len(a[1]) {
			if a[0][i] < a[1][j] {
				s = append(s, a[0][i])
				i++
			} else {
				s = append(s, a[1][j])
				j++
			}
		}
		Fprintf(out, "%s\n", append(append(s, a[0][i:]...), a[1][j:]...))
	}
}

//func main() { CF1251C(os.Stdin, os.Stdout) }
