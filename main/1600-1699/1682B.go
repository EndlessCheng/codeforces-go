package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1682B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		and := -1
		for i := range a {
			Fscan(in, &a[i])
			if a[i] != i {
				and &= a[i]
			}
		}
		Fprintln(out, and)
	}
}

//func main() { cf1682B(bufio.NewReader(os.Stdin), os.Stdout) }
