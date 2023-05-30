package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF846C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, p, q, r int
	var v int64
	Fscan(in, &n)
	s := make([]int64, n+1)
	type pair struct{v int64; i int}
	pre := make([]pair, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s[i+1] = s[i] + v
		if s[i+1] > pre[i].v {
			pre[i+1] = pair{s[i+1], i + 1}
		} else {
			pre[i+1] = pre[i]
		}
	}
	ans := int64(-1e18)
	suf, sufI := s[n], n
	for i := n; i >= 0; i-- {
		if s[i] > suf {
			suf, sufI = s[i], i
		}
		t := pre[i].v - s[i] + suf
		if t > ans {
			ans, p, q, r = t, pre[i].i, i, sufI
		}
	}
	Fprint(out, p, q, r)
}

//func main() { CF846C(os.Stdin, os.Stdout) }
