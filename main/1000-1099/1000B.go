package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1000B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n, m, v, prev, maxDiff int
	s := [2]int{}
	Fscan(in, &n, &m)
	for i := 0; i <= n; i++ {
		if i < n {
			Fscan(in, &v)
		} else {
			v = m
		}
		maxDiff = max(maxDiff, s[0]-s[1])
		s[i&1] += v - prev
		prev = v
	}
	// s[1]+maxDiff-1 相当于 s[1] 去掉前面没转换的 s[1]，同时由于没转换的这部分和 s[0] 的这一部分是绑定在一起的，所以求个循环中的 maxDiff
	Fprint(out, max(s[0], s[1]+maxDiff-1))
}

//func main() { CF1000B(os.Stdin, os.Stdout) }
