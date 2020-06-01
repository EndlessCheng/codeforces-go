package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1363A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, x, v int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &x)
		c := [2]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			c[v&1]++
		}
		if c[1] == 0 || c[0] == 0 && x&1 == 0 || x == n && c[1]&1 == 0 {
			Fprintln(out, "No")
		} else {
			Fprintln(out, "Yes")
		}
	}
}

//func main() { CF1363A(os.Stdin, os.Stdout) }
