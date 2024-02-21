package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf466C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, cnt, ans int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}
	tot := s[n]
	if tot%3 != 0 {
		Fprint(out, 0)
		return
	}
	for i := 1; i < n; i++ {
		if s[i] == tot/3*2 {
			ans += cnt
		}
		if s[i] == tot/3 {
			cnt++
		}
	}
	Fprint(out, ans)
}

//func main() { cf466C(os.Stdin, os.Stdout) }
