package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1030C(_r io.Reader, out io.Writer) {
	var n, s int
	var a []byte
	Fscan(bufio.NewReader(_r), &n, &a)
o:
	for i, v := range a[:n-1] {
		s += int(v & 15)
		for j := i + 1; j < n; {
			s2 := int(a[j] & 15)
			for j++; j < n && (s2 < s || s2 == s && a[j] == '0'); j++ {
				s2 += int(a[j] & 15)
			}
			if s2 != s {
				continue o
			}
		}
		Fprintln(out, "YES")
		return
	}
	Fprintln(out, "NO")
}

func main() { CF1030C(os.Stdin, os.Stdout) }
