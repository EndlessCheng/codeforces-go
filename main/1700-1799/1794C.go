package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf1794C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		l := 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			for a[l] < i-l+1 {
				l++
			}
			Fprint(out, i-l+1, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1794C(bufio.NewReader(os.Stdin), os.Stdout) }
