package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf1101B(in io.Reader, out io.Writer) {
	s := ""
	Fscan(bufio.NewReader(in), &s)
	i := strings.IndexByte(s, '[')
	if i < 0 {
		Fprint(out, -1)
		return
	}
	s = s[i+1:]

	i = strings.IndexByte(s, ':')
	if i < 0 {
		Fprint(out, -1)
		return
	}
	s = s[i+1:]

	i = strings.LastIndexByte(s, ']')
	if i < 0 {
		Fprint(out, -1)
		return
	}
	s = s[:i]

	i = strings.LastIndexByte(s, ':')
	if i < 0 {
		Fprint(out, -1)
		return
	}
	s = s[:i]

	Fprint(out, 4+strings.Count(s, "|"))
}

//func main() { cf1101B(os.Stdin, os.Stdout) }
