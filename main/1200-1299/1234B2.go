package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1234B2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v int
	s := []int{}
	on := map[int]bool{}
	for Fscan(in, &n, &k); n > 0; n-- {
		Fscan(in, &v)
		if !on[v] {
			on[v] = true
			s = append(s, v)
			if len(s) > k {
				delete(on, s[0])
				s = s[1:]
			}
		}
	}
	Fprintln(out, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		Fprint(out, s[i], " ")
	}
}

//func main() { CF1234B2(os.Stdin, os.Stdout) }
