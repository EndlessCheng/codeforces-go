package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1364A(in io.Reader, out io.Writer) {
	var T, n, x, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		s := 0
		l, r := -1, -1
		for i := range n {
			Fscan(in, &v)
			s += v
			if v%x != 0 {
				if l < 0 {
					l = i
				}
				r = i
			}
		}
		if s%x != 0 {
			Fprintln(out, n)
		} else if l < 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, max(n-l-1, r))
		}
	}
}

//func main() { cf1364A(bufio.NewReader(os.Stdin), os.Stdout) }
