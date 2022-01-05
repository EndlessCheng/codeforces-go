package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1594C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s, ch string
O:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &ch, &s)
		if strings.Count(s, ch) == n {
			Fprintln(out, 0)
			continue
		}
	o:
		for i := 1; i <= n; i++ {
			for j := i; j <= n; j += i {
				if s[j-1] != ch[0] {
					continue o
				}
			}
			Fprintln(out, 1)
			Fprintln(out, i)
			continue O
		}
		Fprintln(out, 2)
		Fprintln(out, n, n-1)
	}
}

//func main() { CF1594C(os.Stdin, os.Stdout) }
