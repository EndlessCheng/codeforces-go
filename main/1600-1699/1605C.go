package main

import (
	"bufio"
	. "fmt"
	"io"
	. "strings"
)

// https://space.bilibili.com/206214
func CF1605C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &s)
		ans := -1
		for _, t := range []string{"aa", "aba", "aca", "abca", "acba", "abbacca", "accabba"} {
			if Contains(s, t) {
				ans = len(t)
				break
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1605C(os.Stdin, os.Stdout) }
