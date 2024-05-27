package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf608B(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(in, &s, &t)
	d := len(t) - len(s) + 1
	c1 := strings.Count(t[:d-1], "1")

	ans := 0
	for i, b := range s {
		c1 += int(t[i+d-1] & 1)
		if b == '0' {
			ans += c1
		} else {
			ans += d - c1
		}
		c1 -= int(t[i] & 1)
	}
	Fprint(out, ans)
}

//func main() { cf608B(bufio.NewReader(os.Stdin), os.Stdout) }
