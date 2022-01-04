package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1608B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, u, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &u, &d)
		if u-d > 1 || d-u > 1 || u+d > n-2 {
			Fprintln(out, -1)
			continue
		}
		if u > d {
			v, w := 1, n
			i := 0
			for ; i < n; i++ {
				if i&1 == 0 {
					Fprint(out, v, " ")
					v++
				} else {
					Fprint(out, w, " ")
					w--
					if u--; u == 0 {
						break
					}
				}
			}
			for i++; i < n; i++ {
				Fprint(out, w, " ")
				w--
			}
		} else if u == d {
			v, w := 1, n
			i := 0
			for ; i < n; i++ {
				if i&1 == 0 {
					Fprint(out, v, " ")
					v++
					if d--; d == -1 {
						break
					}
				} else {
					Fprint(out, w, " ")
					w--
				}
			}
			for i++; i < n; i++ {
				Fprint(out, v, " ")
				v++
			}
		} else {
			v, w := 1, n
			i := 0
			for ; i < n; i++ {
				if i&1 == 1 {
					Fprint(out, v, " ")
					v++
					if d--; d == 0 {
						break
					}
				} else {
					Fprint(out, w, " ")
					w--
				}
			}
			for i++; i < n; i++ {
				Fprint(out, v, " ")
				v++
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1608B(os.Stdin, os.Stdout) }
