package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1304B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	reverse := func(a []byte) []byte {
		n := len(a)
		r := make([]byte, n)
		for i, v := range a {
			r[n-1-i] = v
		}
		return r
	}

	ss := [][]byte{}
	pa := []byte{}
	has := map[string]bool{}
	var n, m int
	var s []byte
	Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		if rs := reverse(s); bytes.Equal(s, rs) {
			pa = s
		} else if has[string(rs)] {
			ss = append(ss, s)
		} else {
			has[string(s)] = true
		}
	}

	ans := []byte{}
	for _, s := range ss {
		ans = append(ans, s...)
	}
	rAns := reverse(ans)
	ans = append(ans, pa...)
	ans = append(ans, rAns...)
	Fprintln(out, len(ans))
	Fprint(out, string(ans))
}

//func main() { CF1304B(os.Stdin, os.Stdout) }
