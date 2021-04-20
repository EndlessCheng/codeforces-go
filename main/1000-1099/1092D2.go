package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1092D2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, mx int
	s := []int{}
	for Fscan(in, &n); n > 0; n-- {
		if Fscan(in, &v); v > mx {
			mx = v
		}
		if len(s) == 0 {
			s = append(s, v)
		} else if s[len(s)-1] == v {
			s = s[:len(s)-1]
		} else if s[len(s)-1] < v {
			Fprint(out, "NO")
			return
		} else {
			s = append(s, v)
		}
	}
	if len(s) == 0 || len(s) == 1 && s[0] == mx {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF1092D2(os.Stdin, os.Stdout) }
