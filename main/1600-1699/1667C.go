package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1667C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	k := (n*2 + 1) / 3
	Fprintln(out, k)
	for x, y := 1, 1; y <= k; x, y = x+2, y+1 {
		if x > k {
			x = 2
		}
		Fprintln(out, x, y)
	}
}

//func main() { cf1667C(os.Stdin, os.Stdout) }
