package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1198C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		var e []interface{}
		vis := make([]bool, 3*n+1)
		for i := 1; i <= m; i++ {
			if Fscan(in, &v, &w); len(e) < n && !vis[v] && !vis[w] {
				vis[v] = true
				vis[w] = true
				e = append(e, i)
			}
		}
		if len(e) == n {
			Fprintln(out, "Matching")
			Fprintln(out, e...)
		} else {
			Fprintln(out, "IndSet")
			for i, c := 1, 0; i <= 3*n; i++ {
				if !vis[i] {
					Fprint(out, i, " ")
					if c++; c == n {
						break
					}
				}
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1198C(os.Stdin, os.Stdout) }
