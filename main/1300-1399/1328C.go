package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1328C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		t := []byte{}
		for i, b := range s {
			if b == '1' {
				Fprintf(out, "%s1%s\n", t, strings.Repeat("0", n-1-i))
				Fprintf(out, "%s0%s\n", t, s[i+1:])
				continue o
			}
			t = append(t, '0'+b&15/2)
		}
		Fprintf(out, "%s\n", t)
		Fprintf(out, "%s\n", t)
	}
}

//func main() { CF1328C(os.Stdin, os.Stdout) }
