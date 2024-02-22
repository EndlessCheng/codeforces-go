package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf11A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, d, pre, v, ans int
	for Fscan(in, &n, &d, &pre); n > 1; n-- {
		Fscan(in, &v)
		if v <= pre {
			k := (pre-v)/d + 1
			ans += k
			v += k * d
		}
		pre = v
	}
	Fprint(out, ans)
}

//func main() { cf11A(os.Stdin, os.Stdout) }
