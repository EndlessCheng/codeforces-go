package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1599A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	var s []byte
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	Fscan(in, &s)

	moves := []int{}
	small, big := 0, 1
	for _, b := range s {
		if big&1 > 0 == (b == s[0]) {
			moves = append(moves, big)
			big++
		} else {
			moves = append(moves, small)
			small--
		}
	}

	for _, i := range moves {
		d := "L"
		if i&1 > 0 != (s[0] == 'L') {
			d = "R"
		}
		Fprintln(out, a[i-small-1], d)
	}
}

//func main() { cf1599A(bufio.NewReader(os.Stdin), os.Stdout) }
