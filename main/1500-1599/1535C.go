package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1535C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := int64(0)
		pos := [2]int{-1, -1}
		for i, b := range s {
			if b != '?' {
				pos[(i&1)^int(b&1)] = i
			}
			ans += int64(i - min(pos[0], pos[1]))
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1535C(os.Stdin, os.Stdout) }
