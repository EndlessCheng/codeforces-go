package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1174B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, odd int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		odd += a[i] % 2
	}
	if odd != 0 && odd != n {
		slices.Sort(a)
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { cf1174B(bufio.NewReader(os.Stdin), os.Stdout) }
