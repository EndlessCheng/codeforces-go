package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2020C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, b, c, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &b, &c, &d)
		a := b ^ d
		if a|b-a&c == d {
			Fprintln(out, a)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { cf2020C(bufio.NewReader(os.Stdin), os.Stdout) }
