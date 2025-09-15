package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1174A(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n*2)
	for i := range a {
		Fscan(in, &a[i])
	}

	slices.Sort(a)
	if a[0] == a[n*2-1] {
		Fprint(out, -1)
	} else {
		for _, v := range a {
			Fprint(out, v, " ")
		}
	}
}

//func main() { cf1174A(bufio.NewReader(os.Stdin), os.Stdout) }
