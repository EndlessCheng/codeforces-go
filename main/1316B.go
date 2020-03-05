package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1316B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	reverse := func(ss string) string {
		n := len(ss)
		s := make([]byte, n)
		for i := range s {
			s[i] = ss[n-1-i]
		}
		return string(s)
	}

	var t, n int
	var s, ss string
	Fscan(in, &t)
	for ; t > 0; t-- {
		Fscan(in, &n, &s)
		minS := s
		k := 1
		for i := 1; i < n; i++ {
			if (n-i)%2 == 1 {
				ss = s[i:] + reverse(s[:i])
			} else {
				ss = s[i:] + s[:i]
			}
			if ss < minS {
				minS = ss
				k = i + 1
			}
		}
		Fprintln(out, minS)
		Fprintln(out, k)
	}
}

//func main() { CF1316B(os.Stdin, os.Stdout) }
