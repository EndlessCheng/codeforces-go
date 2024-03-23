package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1443B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, a, b int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &s)
		pre := -9999
		ans := 0
		for i, c := range s {
			if c == '1' {
				ans += min(b*(i-pre-1), a)
				pre = i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1443B(os.Stdin, os.Stdout) }
