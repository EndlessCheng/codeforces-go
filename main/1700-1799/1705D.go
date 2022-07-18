package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1705D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n int
	var s, t string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s, &t)
		if s[0] != t[0] || s[n-1] != t[n-1] {
			Fprintln(out, -1)
			continue
		}
		ans := int64(0)
		i, j := 0, 0
		for i < n && j < n {
			i0 := i
			for i++; i < n && s[i] == s[i-1]; i++ {}
			j0 := j
			for j++; j < n && t[j] == t[j-1]; j++ {}
			if t[j-1] == '1' {
				ans += int64(abs(i0-j0) + abs(i-j))
			}
		}
		if i < n || j < n { ans = -1 }
		Fprintln(out, ans)
	}
}

//func main() { CF1705D(os.Stdin, os.Stdout) }
