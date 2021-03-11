package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1496B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		mx := 0
		s := map[int]bool{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if v > mx {
				mx = v
			}
			s[v] = true
		}
		mex := 0
		for ; s[mex]; mex++ {
		}
		if mex > mx {
			Fprintln(out, len(s)+k)
		} else if k > 0 && !s[(mex+mx+1)/2] {
			Fprintln(out, len(s)+1)
		} else {
			Fprintln(out, len(s))
		}
	}
}

//func main() { CF1496B(os.Stdin, os.Stdout) }
