package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf383A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, cnt1, ans int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if v > 0 {
			cnt1++
		} else {
			ans += cnt1
		}
	}
	Fprint(out, ans)
}

//func main() { cf383A(os.Stdin, os.Stdout) }
