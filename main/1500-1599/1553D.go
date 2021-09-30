package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1553D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		i, j := len(s)-1, len(t)-1
		for i >= 0 && j >= 0 {
			if s[i] == t[j] {
				i--
				j--
			} else {
				i -= 2
			}
		}
		if j < 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1553D(os.Stdin, os.Stdout) }
