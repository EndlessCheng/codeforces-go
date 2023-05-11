package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF1809D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const swap int64 = 1e12
	const rm = swap + 1

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		c0 := strings.Count(s, "0")
		c1 := 0
		ans := int64(min9(c0, len(s)-c0)) * rm
		for i, c := range s {
			if i > 0 && c == '0' && s[i-1] == '1' {
				ans = min9(ans, int64(c0+c1-2)*rm+swap)
			} else {
				ans = min9(ans, int64(c0+c1)*rm)
			}
			c0 -= int(c&1 ^ 1)
			c1 += int(c & 1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1809D(os.Stdin, os.Stdout) }
func min9[T int | int64](a, b T) T {
	if a > b {
		return b
	}
	return a
}
