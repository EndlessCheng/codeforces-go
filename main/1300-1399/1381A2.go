package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1381A2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var a, b []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b)
		ans := []interface{}{}
		for i, c := range a {
			if i > 0 && c != a[i-1] {
				ans = append(ans, i)
			}
		}
		if a[n-1] != b[n-1] {
			ans = append(ans, n)
		}
		for i := len(b) - 2; i >= 0; i-- {
			if b[i] != b[i+1] {
				ans = append(ans, i+1)
			}
		}
		Fprint(out, len(ans), " ")
		Fprintln(out, ans...)
	}
}

//func main() { CF1381A2(os.Stdin, os.Stdout) }
