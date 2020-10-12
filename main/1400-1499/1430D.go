package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1430D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		c := []int{}
		for i := 0; i < n; {
			st := i
			for ; i < n && s[i] == s[st]; i++ {
			}
			c = append(c, i-st)
		}
		n = len(c)
		ans := 0
		for i, p := 0, 0; i < n; i++ {
			if p < i {
				p = i
			}
			for ; p < n && c[p] <= 1; p++ {
			}
			if p < n {
				c[p]--
			} else {
				i++
			}
			ans++
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1430D(os.Stdin, os.Stdout) }
