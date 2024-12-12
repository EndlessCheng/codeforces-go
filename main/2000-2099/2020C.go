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
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &b, &c, &d)
		a := 0
		for bit := 1 << 59; bit > 0; bit >>= 1 {
			if d&bit == 0 {
				if b&bit > 0 && c&bit == 0 {
					Fprintln(out, -1)
					continue o
				}
				a |= c & bit
			} else if b&bit == 0 {
				if c&bit > 0 {
					Fprintln(out, -1)
					continue o
				}
				a |= bit
			}
		}
		Fprintln(out, a)
	}
}

//func main() { cf2020C(bufio.NewReader(os.Stdin), os.Stdout) }
