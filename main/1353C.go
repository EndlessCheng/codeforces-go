package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1353C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	a := [5e5]int64{}
	for i := int64(3); i < 5e5; i += 2 {
		a[i] = a[i-2] + i/2*4*(i-1)
	}

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		Fprintln(out, a[n])
	}
}

//func main() { CF1353C(os.Stdin, os.Stdout) }
