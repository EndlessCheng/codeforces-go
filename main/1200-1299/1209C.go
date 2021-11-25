package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1209C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
O:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
	o:
		for up := '0'; up <= '9'; up++ {
			ans := bytes.Repeat([]byte{'2'}, n)
			var pre1, pre2 rune
			for i, b := range s {
				if b == up && pre2 == 0 {
					continue
				}
				if b <= up {
					if b < pre1 {
						continue o
					}
					ans[i] = '1'
					pre1 = b
				} else {
					if b < pre2 {
						continue o
					}
					pre2 = b
				}
			}
			Fprintf(out, "%s\n", ans)
			continue O
		}
		Fprintln(out, "-")
	}
}

//func main() { CF1209C(os.Stdin, os.Stdout) }
