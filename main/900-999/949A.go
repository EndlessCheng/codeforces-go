package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF949A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(bufio.NewReader(in), &s)
	a, p := [][]int{}, 0
	for i, b := range s {
		if b == '0' {
			if p == len(a) {
				a = append(a, []int{i})
			} else {
				a[p] = append(a[p], i)
			}
			p++
		} else {
			if p == 0 {
				Fprint(out, -1)
				return
			}
			p--
			a[p] = append(a[p], i)
		}
	}
	if p != len(a) {
		Fprint(out, -1)
		return
	}

	Fprintln(out, p)
	for _, b := range a {
		Fprint(out, len(b))
		for _, i := range b {
			Fprint(out, " ", i+1)
		}
		Fprintln(out)
	}
}

//func main() { CF949A(os.Stdin, os.Stdout) }
