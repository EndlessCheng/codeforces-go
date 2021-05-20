package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF883F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s string
	mp := map[string]bool{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		for {
			t := strings.ReplaceAll(s, "kh", "h")
			if t == s {
				break
			}
			s = t
		}
		mp[strings.ReplaceAll(s, "u", "oo")] = true
	}
	Fprint(out, len(mp))
}

//func main() { CF883F(os.Stdin, os.Stdout) }
