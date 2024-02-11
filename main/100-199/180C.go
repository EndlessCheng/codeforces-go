package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf180C(in io.Reader, out io.Writer) {
	var s string
	Fscan(bufio.NewReader(in), &s)
	f, cnt := 0, 0
	for _, b := range s {
		if b >= 'a' {
			cnt++
		} else {
			f = min(f+1, cnt)
		}
	}
	Fprint(out, f)
}

//func main() { cf180C(os.Stdin, os.Stdout) }
