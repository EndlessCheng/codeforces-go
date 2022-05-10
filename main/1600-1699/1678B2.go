package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1678B2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans, seg := 0, 0
		for i, pre := 0, byte(0); i < n; i += 2 {
			if s[i] != s[i+1] {
				ans++
			} else if s[i] != pre {
				seg++
				pre = s[i]
			}
		}
		if seg == 0 {
			seg = 1
		}
		Fprintln(out, ans, seg)
	}
}

//func main() { CF1678B2(os.Stdin, os.Stdout) }
