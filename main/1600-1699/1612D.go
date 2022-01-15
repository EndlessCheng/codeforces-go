package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1612D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b, x int64
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &x)
		if a > b {
			a, b = b, a
		}
		for a > 0 && x <= b {
			if x >= a && (b-x)%a == 0 {
				Fprintln(out, "YES")
				continue o
			}
			a, b = b%a, a
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1612D(os.Stdin, os.Stdout) }
