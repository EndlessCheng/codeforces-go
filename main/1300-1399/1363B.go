package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1363B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var t int
	var S []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &S)
		n := len(S)
		s := make([][2]int, n+1)
		for i, b := range S {
			s[i+1] = s[i]
			s[i+1][b-'0']++
		}
		ans := n
		for _, p := range s {
			ans = min(ans, min(p[1]+s[n][0]-p[0], p[0]+s[n][1]-p[1]))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1363B(os.Stdin, os.Stdout) }
