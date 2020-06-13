package main

import (
	"bufio"
	. "fmt"
	"io"
	. "strings"
)

// github.com/EndlessCheng/codeforces-go
func CF691C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	if !Contains(s, ".") {
		s += "."
	}
	s = Trim(s, "0")

	i := Index(s, ".")
	e := i - IndexAny(s, "123456789")
	if e > 0 {
		e--
	}
	s = Trim(s[:i]+s[i+1:], "0")
	if len(s) > 1 {
		s = s[:1] + "." + s[1:]
	}
	Fprint(out, s)
	if e != 0 {
		Fprint(out, "E", e)
	}
}

//func main() { CF691C(os.Stdin, os.Stdout) }
