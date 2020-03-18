package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1307B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, x, v int
	for Fscan(in, &t); t > 0; t-- {
		mx := 0
		found := false
		for Fscan(in, &n, &x); n > 0; n-- {
			if Fscan(in, &v); v > mx {
				mx = v
			}
			if v == x {
				found = true
			}
		}
		if found {
			Fprintln(out, 1)
		} else if mx > x {
			Fprintln(out, 2)
		} else {
			Fprintln(out, (x-1)/mx+1)
		}
	}
}

//func main() { CF1307B(os.Stdin, os.Stdout) }
