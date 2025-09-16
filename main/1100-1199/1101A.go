package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1101A(in io.Reader, out io.Writer) {
	var T, l, r, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r, &d)
		if d >= l {
			d *= r/d + 1
		}
		Fprintln(out, d)
	}
}

//func main() { cf1101A(bufio.NewReader(os.Stdin), os.Stdout) }
