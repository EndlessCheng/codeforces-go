package main

import (
	"bufio"
	. "fmt"
	"io"
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF490E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	nxt := func(s string) string { v, _ := strconv.Atoi(s); return strconv.Itoa(v + 1) }

	var n int
	Fscan(in, &n)
	pre := ""
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
		s := a[i]
		if len(s) < len(pre) {
			Fprint(out, "NO")
			return
		}
		if len(s) > len(pre) {
			a[i] = strings.ReplaceAll(s[:1], "?", "1") + strings.ReplaceAll(s[1:], "?", "0")
			pre = nxt(a[i])
			continue
		}
		j := 0
	o:
		for ; j < len(s); j++ {
			b := s[j]
			if b == '?' || b == pre[j] {
				continue
			}
			if b > pre[j] {
				a[i] = pre[:j] + strings.ReplaceAll(s[j:], "?", "0")
				break
			}
			for j--; j >= 0; j-- {
				if s[j] == '?' && pre[j] < '9' {
					a[i] = pre[:j] + string(pre[j]+1) + strings.ReplaceAll(s[j+1:], "?", "0")
					break o
				}
			}
			Fprint(out, "NO")
			return
		}
		if j == len(s) {
			a[i] = pre
		}
		pre = nxt(a[i])
	}
	Fprintln(out, "YES")
	for _, s := range a {
		Fprintln(out, s)
	}
}

//func main() { CF490E(os.Stdin, os.Stdout) }
