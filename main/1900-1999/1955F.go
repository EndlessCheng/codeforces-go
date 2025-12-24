package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1955F(in io.Reader, out io.Writer) {
	var T, a, b, c, d int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c, &d)
		Fprintln(out, a/2+b/2+c/2+d/2+a*b*c%2)
	}
}

//func main() { cf1955F(bufio.NewReader(os.Stdin), os.Stdout) }
