package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1392D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		i := strings.Index(s, "LR")
		if i < 0 {
			i = strings.Index(s, "RL")
			if i < 0 {
				Fprintln(out, (n+2)/3)
				continue
			}
		}
		s = s[i+1:] + s[:i+1]
		ans := 0
		for i := 0; i < n; {
			st := i
			for ; i < n && s[i] == s[st]; i++ {
			}
			ans += (i - st) / 3
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1392D(os.Stdin, os.Stdout) }
