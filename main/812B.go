package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF812B(_r io.Reader, _w io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, ansL, ansR, posL, posR, maxFloor int
	Fscan(in, &n, &m)
	m++
	floors := make([]string, n)
	for i := range floors {
		Fscan(in, &floors[n-1-i])
	}
	for i, s := range floors {
		l := strings.IndexByte(s, '1')
		if l == -1 {
			continue
		}
		maxFloor = i
		r := strings.LastIndexByte(s, '1')
		if ansL == 0 {
			ansL, ansR = r, r
			posL, posR = r, r
		} else {
			ansL, ansR = min(ansL+m-posL, ansR+m-posR)+m-l, min(ansL+posL, ansR+posR)+r
			posL, posR = l, r
		}
	}
	Fprint(out, min(ansL, ansR)+maxFloor)
}

//func main() {
//	CF812B(os.Stdin, os.Stdout)
//}
