package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF729D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, a, b, k int
	var s string
	Fscan(in, &n, &a, &b, &k, &s)
	pos := []int{}
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			continue
		}
		st := i
		for i++; i < n && s[i] == '0'; i++ {
		}
		for st += b - 1; st < i; st += b {
			pos = append(pos, st)
		}
	}
	pos = pos[a-1:]
	Fprintln(out, len(pos))
	for _, v := range pos {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF729D(os.Stdin, os.Stdout) }
