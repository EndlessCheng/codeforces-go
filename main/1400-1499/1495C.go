package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1495C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]byte, n)
		for i := range a {
			Fscan(in, &a[i])
			if i%3 == (3-n%3)/3 {
				a[i] = bytes.Repeat([]byte{'X'}, m)
				j := 0
				if m > 1 && (i > 0 && a[i-1][1] == 'X' || i > 1 && a[i-2][1] == 'X') {
					j = 1
				}
				if i > 0 {
					a[i-1][j] = 'X'
				}
				if i > 1 {
					a[i-2][j] = 'X'
				}
			}
		}
		for _, r := range a {
			Fprintln(out, string(r))
		}
	}
}

//func main() { CF1495C(os.Stdin, os.Stdout) }
