package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1708B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, l, r int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		a := make([]int, n)
		for i := range a {
			a[i] = r - r%(i+1)
			if a[i] < l {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1708B(bufio.NewReader(os.Stdin), os.Stdout) }
