package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1157B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var s []byte
	Fscan(in, &n, &s)
	mp := [10]byte{}
	for i := 1; i < 10; i++ {
		Fscan(in, &mp[i])
	}
	start := false
	for i, b := range s {
		b &= 15
		if start && mp[b] < b {
			break
		}
		if mp[b] > b {
			s[i] = '0' + mp[b]
			start = true
		}
	}
	Fprintf(out, "%s", s)
}

//func main() { CF1157B(os.Stdin, os.Stdout) }
