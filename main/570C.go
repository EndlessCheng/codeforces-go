package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF570C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, p, dot, seg int
	var s, c []byte
	Fscan(in, &n, &q, &s)
	for i, b := range s {
		if b == '.' {
			dot++
			if i == 0 || s[i-1] != '.' {
				seg++
			}
		}
	}
	s = append([]byte{0}, append(s, 0)...)
	for ; q > 0; q-- {
		Fscan(in, &p, &c)
		if s[p] == '.' != (c[0] == '.') {
			d := -1
			if c[0] == '.' {
				d = 1
			}
			dot += d
			if s[p-1] != '.' && s[p+1] != '.' {
				seg += d
			} else if s[p-1] == '.' && s[p+1] == '.' {
				seg -= d
			}
			s[p] = c[0]
		}
		Fprintln(out, dot-seg)
	}
}

//func main() { CF570C(os.Stdin, os.Stdout) }
