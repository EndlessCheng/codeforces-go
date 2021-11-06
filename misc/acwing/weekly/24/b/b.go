package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(in, &s, &t)
	x0, y0 := int(s[0]-'a'), int(s[1]-'1')
	x1, y1 := int(t[0]-'a'), int(t[1]-'1')
	ans := 0
	for x2 := 0; x2 < 8; x2++ {
		for y2 := 0; y2 < 8; y2++ {
			if x2 != x0 && y2 != y0 && abs((x2-x0)*(y2-y0)) != 2 && abs((x2-x1)*(y2-y1)) != 2 {
				ans++
			}
		}
	}
	Fprint(out, ans-1)
}

func main() { run(os.Stdin, os.Stdout) }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
