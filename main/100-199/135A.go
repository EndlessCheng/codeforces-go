package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

func cf135A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, mxI int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] > a[mxI] {
			mxI = i
		}
	}
	if a[mxI] == 1 {
		a[mxI] = 2
	} else {
		a[mxI] = 1
	}
	slices.Sort(a)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { cf135A(bufio.NewReader(os.Stdin), os.Stdout) }
