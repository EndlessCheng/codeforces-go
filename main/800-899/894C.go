package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF894C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i]%a[0] != 0 {
			Fprint(out, -1)
			return
		}
	}
	Fprintln(out, 2*n)
	for _, v := range a {
		Fprint(out, v, a[0], " ")
	}
}

//func main() { CF894C(os.Stdin, os.Stdout) }
