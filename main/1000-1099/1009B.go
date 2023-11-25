package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF1009B(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	s1 := strings.Repeat("1", strings.Count(s, "1"))
	s = strings.ReplaceAll(s, "1", "")
	i := strings.IndexByte(s, '2')
	if i < 0 {
		Fprint(out, s, s1)
	} else {
		Fprint(out, s[:i], s1, s[i:])
	}
}

//func main() { CF1009B(os.Stdin, os.Stdout) }
