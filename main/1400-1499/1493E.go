package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1493E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var l, r string
	Fscan(in, &n, &l, &r)
	if l[0] != r[0] {
		Fprint(out, strings.Repeat("1", n))
		return
	}
	// 判断 r-l > 1 的非 bigint 写法
	// 判断方法是找到第一个不同的位，并且后续不是 l=011...1, r=100...0 的情况
	// 此时 r-l > 1 成立
	i := 0
	for ; i < n; i++ {
		if r[i] == '1' && l[i] == '0' {
			break
		}
	}
	for i++; i < n; i++ {
		if r[i] == '1' || l[i] == '0' {
			Fprint(out, r[:n-1]+"1")
			return
		}
	}
	Fprint(out, r)
}

//func main() { CF1493E(os.Stdin, os.Stdout) }
