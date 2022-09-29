package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1178E(in io.Reader, out io.Writer) {
	var s, ans []byte
	Fscan(bufio.NewReader(in), &s)
	n := len(s)
	for i := 1; i < n/2; i += 2 {
		if s[i] == s[n-i] || s[i] == s[n-i-1] {
			ans = append(ans, s[i])
		} else {
			ans = append(ans, s[i-1])
		}
	}
	Fprintf(out, "%s", ans)
	if n%4 > 0 {
		Fprintf(out, "%c", s[n/2])
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF1178E(os.Stdin, os.Stdout) }
