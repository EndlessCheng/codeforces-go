package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1675E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		if k > 24 {
			Fprintln(out, strings.Repeat("a", n))
			continue
		}
		var mx, up, low byte
		mx = 'a'
		for _, b := range s {
			if b > mx {
				if k := byte(k) - (mx - 'a'); b-k > mx {
					up, low = b, b-k
					break
				}
				mx = b
			}
		}
		for i, b := range s {
			if b <= mx {
				s[i] = 'a'
			} else if low <= b && b <= up {
				s[i] = low
			}
		}
		Fprintf(out, "%s\n", s)
	}
}

//func main() { CF1675E(os.Stdin, os.Stdout) }
