package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1385D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var f func(s string, b byte) int
	f = func(s string, b byte) int {
		n := len(s)
		if n == 1 {
			if s[0] == b {
				return 0
			}
			return 1
		}
		n >>= 1
		return n + min(f(s[n:], b+1)-strings.Count(s[:n], string(b)), f(s[:n], b+1)-strings.Count(s[n:], string(b)))
	}

	var t, n int
	var s string
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &s)
		Fprintln(out, f(s, 'a'))
	}
}

//func main() { CF1385D(os.Stdin, os.Stdout) }
