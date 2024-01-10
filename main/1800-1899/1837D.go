package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf1837D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, n, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := make([]any, n)
		cnt := 0
		mask := 0
		for i, b := range s {
			if b == '(' {
				cnt++
			} else {
				cnt--
			}
			if cnt > 0 {
				ans[i] = 1
				mask |= 1
			} else if cnt < 0 {
				ans[i] = 2
				mask |= 2
			} else {
				ans[i] = ans[i-1]
			}
		}
		if cnt != 0 {
			Fprintln(out, -1)
		} else if mask < 3 {
			Fprintln(out, 1)
			Fprintln(out, strings.Repeat("1 ", n))
		} else {
			Fprintln(out, 2)
			Fprintln(out, ans...)
		}
	}
}

//func main() { cf1837D(os.Stdin, os.Stdout) }
