package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1504A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		if i := strings.IndexFunc(s, func(r rune) bool { return r > 'a' }); i < 0 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			Fprintln(out, s[:len(s)-i]+"a"+s[len(s)-i:])
		}
	}
}

//func main() { CF1504A(os.Stdin, os.Stdout) }
