package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF978C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	var v int64
	Fscan(in, &n, &m)
	s := make([]int64, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s[i+1] = s[i] + v
	}
	for i := 1; m > 0; m-- {
		for Fscan(in, &v); s[i] < v; i++ {
		}
		Fprintln(out, i, v-s[i-1])
	}
}

//func main() { CF978C(os.Stdin, os.Stdout) }
