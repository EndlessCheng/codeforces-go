package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1579E1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		var a, b []int
		for Fscan(in, &n); n > 0; n-- {
			if Fscan(in, &v); a == nil || v < a[len(a)-1] {
				a = append(a, v)
			} else {
				b = append(b, v)
			}
		}
		for i := len(a) - 1; i >= 0; i-- {
			Fprint(out, a[i], " ")
		}
		for _, v := range b {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1579E1(os.Stdin, os.Stdout) }
