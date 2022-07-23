package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF725C(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	pos := map[byte]int{}
	for i, c := range s {
		p, ok := pos[c]
		if !ok {
			pos[c] = i
			continue
		}
		d := i - p
		if d == 1 {
			Fprint(out, "Impossible")
		} else {
			s = append(append(s[p+d/2+1:], s[:p]...), s[p+1:p+d/2+1]...)
			for i := 0; i < 6; i++ {
				s[i], s[12-i] = s[12-i], s[i]
			}
			Fprintf(out, "%s\n%s", s[:13], s[13:])
		}
		return
	}
}

//func main() { CF725C(os.Stdin, os.Stdout) }
