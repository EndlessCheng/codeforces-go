package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1348C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var b []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &b)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		if s := string(b); s[k-1] != s[0] {
			Fprintln(out, string(s[k-1]))
		} else if k == n || s[k] != s[n-1] {
			Fprintln(out, s[k-1:])
		} else {
			Fprintln(out, s[:1]+strings.Repeat(s[n-1:], (n-1)/k))
		}
	}
}

//func main() { CF1348C(os.Stdin, os.Stdout) }
