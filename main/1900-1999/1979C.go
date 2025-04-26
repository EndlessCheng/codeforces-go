package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1979C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			a[i] = 232792560 / a[i]
			s += a[i]
		}
		if s >= 232792560 {
			Fprintln(out, -1)
		} else {
			for _, v := range a {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf1979C(bufio.NewReader(os.Stdin), os.Stdout) }
