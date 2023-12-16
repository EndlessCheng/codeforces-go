package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf1251B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	T, n, s := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		hasOdd := false
		c1 := 0
		for i := 0; i < n; i++ {
			Fscan(in, &s)
			if len(s)%2 > 0 {
				hasOdd = true
			}
			c1 += strings.Count(s, "1")
		}
		ans := n
		if !hasOdd && c1%2 > 0 {
			ans--
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1251B(os.Stdin, os.Stdout) }
