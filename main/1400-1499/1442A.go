package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1442A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &p)
		s := p
		for n--; n > 0; n-- {
			Fscan(in, &v)
			if v < p {
				s -= p - v
			}
			p = v
		}
		if s < 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1442A(os.Stdin, os.Stdout) }
