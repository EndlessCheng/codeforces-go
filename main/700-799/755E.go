package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf755E(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	ex := 4
	if m == 2 {
		ex++
	}
	if m < 2 || m > 3 || n < ex {
		Fprint(out, -1)
		return
	}
	Fprintln(out, n-1)
	for i := 2; i <= n; i++ {
		x := 3
		if m == 2 || i <= 4 {
			x = i - 1
		}
		Fprintln(out, x, i)
	}
}

//func main() { cf755E(bufio.NewReader(os.Stdin), os.Stdout) }
