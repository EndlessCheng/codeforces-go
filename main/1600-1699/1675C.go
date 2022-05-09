package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1675C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		i, j := strings.LastIndex(s, "1"), strings.Index(s, "0")
		if i < 0 {
			i = 0
		}
		if j < 0 {
			j = len(s) - 1
		}
		Fprintln(out, j-i+1)
	}
}

//func main() { CF1675C(os.Stdin, os.Stdout) }
