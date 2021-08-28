package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

/*
用 KMP 可以求出既是 s 前缀也是 s 后缀的子串长度 match
统计出 match[1..n-1] 中的最大值 mxM
若 t 在 s 中间存在，则 match[n-1], match[match[n-1]-1], ... 中必然存在一个不超过 mxM 的数
*/

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		match := make([]int, n)
		mxM := 0
		for i, c := 1, 0; i < n; i++ {
			b := s[i]
			for c > 0 && s[c] != b {
				c = match[c-1]
			}
			if s[c] == b {
				c++
			}
			match[i] = c
			if i < n-1 && c > mxM {
				mxM = c
			}
		}
		for c := match[n-1]; c > 0; c = match[c-1] {
			if c <= mxM {
				Fprintln(out, s[:c])
				continue o
			}
		}
		Fprintln(out, "not exist")
	}
}

func main() { run(os.Stdin, os.Stdout) }
