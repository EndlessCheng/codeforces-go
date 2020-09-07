package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1404A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		for i, b := range s[k:] {
			if b != '?' {
				if s[i%k] == '?' {
					s[i%k] = b
				} else if s[i%k] != b {
					Fprintln(out, "NO")
					continue o
				}
			}
		}
		if bytes.Count(s[:k], []byte{'0'}) > k/2 || bytes.Count(s[:k], []byte{'1'}) > k/2 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1404A(os.Stdin, os.Stdout) }
