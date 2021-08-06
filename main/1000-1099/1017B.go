package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1017B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var a, b string
	Fscan(in, &n, &a, &b)
	cnt := [2][2]int64{}
	for i, c := range a {
		cnt[c&1][b[i]&1]++
	}
	Fprint(out, cnt[1][0]*cnt[0][1]+cnt[0][0]*(cnt[1][1]+cnt[1][0]))
}

//func main() { CF1017B(os.Stdin, os.Stdout) }
