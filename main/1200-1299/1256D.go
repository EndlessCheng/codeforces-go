package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func CF1256D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var k int64
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		c0, c1 := 0, 0
		for i, b := range s {
			if b == '0' {
				if k < int64(c1) {
					Fprintln(out, strings.Repeat("0", c0)+strings.Repeat("1", c1-int(k))+"0"+strings.Repeat("1", int(k))+s[i+1:])
					continue o
				}
				k -= int64(c1)
				c0++
			} else {
				c1++
			}
		}
		Fprintln(out, strings.Repeat("0", c0)+strings.Repeat("1", c1))
	}
}

//func main() { CF1256D(os.Stdin, os.Stdout) }
