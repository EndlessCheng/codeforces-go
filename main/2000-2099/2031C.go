package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2031C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	a := []any{1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 1, 6}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%2 == 0 {
			for i := 1; i <= n/2; i++ {
				Fprint(out, i, " ", i, " ")
			}
		} else if n < 27 {
			Fprint(out, -1)
		} else {
			Print(a...)
			n -= 27
			for i := 14; i < 14+n/2; i++ {
				Fprint(out, " ", i, " ", i)
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2031C(bufio.NewReader(os.Stdin), os.Stdout) }
