package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF526D(in io.Reader, out io.Writer) {
	var n, k, cnt int
	var s string
	Fscan(bufio.NewReader(in), &n, &k, &s)
	if k == 1 {
		Fprint(out, strings.Repeat("1", n))
		return
	}
	pi := make([]int, n)
	ans := bytes.Repeat([]byte{'0'}, n)
	for i := 1; i < n; i++ {
		b := s[i]
		for cnt > 0 && s[cnt] != b {
			cnt = pi[cnt-1]
		}
		if s[cnt] == b {
			cnt++
		}
		pi[i] = cnt
		if cnt == 0 {
			continue
		}
		total := i + 1
		period := total - cnt
		ab := total / period / k * period
		a := total - ab*k
		if ab >= a {
			ans[i] = '1'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF526D(os.Stdin, os.Stdout) }
