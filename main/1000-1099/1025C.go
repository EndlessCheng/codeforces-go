package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1025C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s string
	Fscan(in, &s)
	n := len(s)
	s += s
	ans := 0
	for i := 0; i < n; {
		st := i
		for i++; i < n*2 && s[i] != s[i-1]; i++ {
		}
		if i-st > ans {
			ans = i - st
		}
	}
	if ans > n {
		ans = n
	}
	Fprint(out, ans)
}

//func main() { CF1025C(os.Stdin, os.Stdout) }
