package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1437B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		c, c2 := 0, 0
		for _, s := range strings.FieldsFunc(s, func(r rune) bool { return r == '0' }) {
			c += len(s) - 1
		}
		for _, s := range strings.FieldsFunc(s, func(r rune) bool { return r == '1' }) {
			c2 += len(s) - 1
		}
		if c2 > c {
			c = c2
		}
		Fprintln(out, c)
	}
}

//func main() { CF1437B(os.Stdin, os.Stdout) }
