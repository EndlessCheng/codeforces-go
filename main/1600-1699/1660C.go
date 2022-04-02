package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1660C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := len(s)
		v := 0
		for _, b := range s {
			b := 1 << (b & 31)
			if v&b > 0 {
				ans -= 2
				v = 0
			} else {
				v |= b
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1660C(os.Stdin, os.Stdout) }
