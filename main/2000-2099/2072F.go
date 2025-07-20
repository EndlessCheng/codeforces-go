package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2072F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	T, n, k := 0, 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		k += " "
		for i := range n {
			if (n-1)&i == i {
				Fprint(out, k)
			} else {
				Fprint(out, "0 ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2072F(bufio.NewReader(os.Stdin), os.Stdout) }
