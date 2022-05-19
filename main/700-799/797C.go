package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF797C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s []byte
	Fscan(in, &s)
	st := s[:0]
	ans := make([]byte, 0, len(s))
	cnt := ['z' + 1]int{}
	for _, b := range s {
		cnt[b]++
	}
	nxt := byte('a')
	for _, b := range s {
		cnt[b]--
		for nxt < 'z' && cnt[nxt] == 0 {
			nxt++
		}
		st = append(st, b)
		for len(st) > 0 && st[len(st)-1] <= nxt {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF797C(os.Stdin, os.Stdout) }
