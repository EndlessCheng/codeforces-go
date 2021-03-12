package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF675C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	c := map[int64]int{}
	var n, mx int
	var s, v int64
	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		s += v
		if c[s]++; c[s] > mx {
			mx = c[s]
		}
	}
	Fprint(out, n-mx)
}

//func main() { CF675C(os.Stdin, os.Stdout) }
