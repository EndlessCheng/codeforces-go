package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF33C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, v, tot, s, maxS int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		tot += v
		s = max(s+v, v)
		maxS = max(maxS, s)
	}
	Fprint(out, 2*maxS-tot)
}

//func main() { CF33C(os.Stdin, os.Stdout) }
