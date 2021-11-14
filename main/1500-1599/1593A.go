package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1593A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	a := [3]int{}
	for Fscan(in, &T); T > 0; T-- {
		mx, c := -1, 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] > mx {
				mx, c = a[i], 1
			} else if a[i] == mx {
				c++
			}
		}
		for _, v := range a {
			if v == mx {
				if c == 1 {
					Fprint(out, "0 ")
				} else {
					Fprint(out, "1 ")
				}
			} else {
				Fprint(out, mx-v+1, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1593A(os.Stdin, os.Stdout) }
