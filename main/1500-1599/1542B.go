package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1542B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var n, a, b int64
	f := func() bool {
		if a == 1 {
			return (n-1)%b == 0
		}
		for i := int64(1); i <= n; i *= a {
			if (n-i)%b == 0 {
				return true
			}
		}
		return false
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b)
		if f() {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { CF1542B(os.Stdin, os.Stdout) }
