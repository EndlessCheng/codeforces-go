package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf877C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)

	Fprintln(out, n+n/2)
	for i := 2; i <= n; i += 2 {
		Fprint(out, i, " ")
	}
	for i := 1; i <= n; i += 2 {
		Fprint(out, i, " ")
	}
	for i := 2; i <= n; i += 2 {
		Fprint(out, i, " ")
	}
}

//func main() { cf877C(bufio.NewReader(os.Stdin), os.Stdout) }
