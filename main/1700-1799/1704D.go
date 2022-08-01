package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1704D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	var v int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		for i, pre := 1, int64(0); i <= n; i++ {
			s := int64(0)
			for j := 0; j < m; j++ {
				Fscan(in, &v)
				s += int64(j) * v
			}
			if pre == -1 {
				continue
			}
			if pre == 0 {
				pre = s
			} else if pre != s {
				if s > pre {
					Fprintln(out, i, s-pre)
				} else {
					Fprintln(out, i-1, pre-s)
				}
				pre = -1
			}
		}
	}
}

//func main() { CF1704D(os.Stdin, os.Stdout) }
