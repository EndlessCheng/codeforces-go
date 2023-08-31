package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ba, a, b, ans int
	var s string
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		ans += strings.Count(s, "AB")
		if s[0] == 'B' && s[len(s)-1] == 'A' {
			ba++
		} else if s[len(s)-1] == 'A' {
			a++
		} else if s[0] == 'B' {
			b++
		}
	}
	if ba > 0 && a == 0 && b == 0 {
		ans--
	}
	Fprint(out, ans+ba+min(a, b))
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
