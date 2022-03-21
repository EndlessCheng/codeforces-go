package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1654B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		pos := [26]int{}
		for i := len(s) - 1; i >= 0; i-- {
			if b := s[i] - 'a'; pos[b] == 0 {
				pos[b] = i + 1
			}
		}
		minP := len(s)
		for _, p := range pos {
			if p > 0 && p < minP {
				minP = p
			}
		}
		Fprintln(out, s[minP-1:])
	}
}

//func main() { CF1654B(os.Stdin, os.Stdout) }
