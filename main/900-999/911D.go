package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf911D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, inv, m, l, r int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		for _, v := range a[:i] {
			if v > a[i] {
				inv ^= 1
			}
		}
	}

	Fscan(in, &m)
	for range m {
		Fscan(in, &l, &r)
		sz := r - l + 1
		inv ^= sz * (sz - 1) / 2 % 2
		if inv == 0 {
			Fprintln(out, "even")
		} else {
			Fprintln(out, "odd")
		}
	}
}

//func main() { cf911D(bufio.NewReader(os.Stdin), os.Stdout) }
