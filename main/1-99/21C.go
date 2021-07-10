package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF21C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, 0)
		return
	}

	a := make([]int, n)
	s := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		s[i+1] = s[i] + a[i]
	}
	if s[n]%3 != 0 {
		Fprint(out, 0)
		return
	}

	s3 := s[n] / 3
	c2 := 0
	for _, v := range s[1:n] {
		if v == s3*2 {
			c2++
		}
	}
	ans := int64(0)
	for _, v := range s[1:n] {
		if v == s3*2 {
			c2--
		}
		if v == s3 {
			ans += int64(c2)
		}
	}
	Fprint(out, ans)
}

//func main() { CF21C(os.Stdin, os.Stdout) }
