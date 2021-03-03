package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1494B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	var a [4]int
T:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &a[0], &a[1], &a[2], &a[3])
	o:
		for i := 0; i < 16; i++ {
			x := i | i&1<<4
			for j, v := range a {
				if v == 0 && x>>j&3 > 0 || v == 1 && x>>j&3 == 3 || v == n-1 && x>>j&3 == 0 || v == n && x>>j&3 < 3 {
					continue o
				}
			}
			Fprintln(out, "YES")
			continue T
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1494B(os.Stdin, os.Stdout) }
