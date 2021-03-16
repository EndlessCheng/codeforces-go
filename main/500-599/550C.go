package main

import (
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF550C(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	if strings.Contains(s, "0") {
		Fprint(out, "YES\n0")
		return
	}
	if strings.Contains(s, "8") {
		Fprint(out, "YES\n8")
		return
	}
	for i := range s {
		v := s[i]
		for j := i + 1; j < len(s); j++ {
			w := s[j]
			if (v&15*10+w&15)%8 == 0 {
				Fprintf(out, "YES\n%c%c", v, w)
				return
			}
			for k := j + 1; k < len(s); k++ {
				if x := s[k]; (v&15*100+w&15*10+x&15)%8 == 0 {
					Fprintf(out, "YES\n%c%c%c", v, w, x)
					return
				}
			}
		}
	}
	Fprint(out, "NO")
}

//func main() { CF550C(os.Stdin, os.Stdout) }
