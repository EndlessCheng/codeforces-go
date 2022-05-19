package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF797C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s, ans, st []byte
	Fscan(in, &s)
	cnt := ['z' + 1]int{}
	for _, b := range s {
		cnt[b]++
	}
	nxt := byte('a')
	for _, b := range s {
		for cnt[nxt] == 0 {
			nxt++
		}
		cnt[b]--
		for len(st) > 0 && st[len(st)-1] <= nxt {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
		st = append(st, b)
	}
	for i := len(st) - 1; i >= 0; i-- {
		ans = append(ans, st[i])
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF797C(os.Stdin, os.Stdout) }
