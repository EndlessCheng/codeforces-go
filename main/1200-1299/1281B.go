package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1281B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		last := [26]int{}
		for i := range last {
			last[i] = -1
		}
		for i, b := range s {
			last[b-'A'] = i
		}
		k := byte(0)
		for i, b := range s {
			for k < b-'A' && last[k] < i {
				k++
			}
			if k < b-'A' {
				s[i], s[last[k]] = s[last[k]], s[i]
				break
			}
		}
		if bytes.Compare(s, t) < 0 {
			Fprintf(out, "%s\n", s)
		} else {
			Fprintln(out, "---")
		}
	}
}

//func main() { CF1281B(os.Stdin, os.Stdout) }
