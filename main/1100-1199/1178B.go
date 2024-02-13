package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1178B(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	var f0, f1, f2 int
	for i := 1; i < len(s); i++ {
		if s[i] == 'o' {
			f1 += f0
		} else if s[i-1] == 'v' {
			f2 += f1
			f0++
		}
	}
	Fprint(out, f2)
}

//func main() { cf1178B(os.Stdin, os.Stdout) }
