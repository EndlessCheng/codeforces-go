package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1918G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	if n <= 5 && n&1 != 0 {
		Fprintln(out, "NO")
		return
	}
	Fprintln(out, "YES")
	if n&1 != 0 {
		Fprint(out, "5 -3 -2 1 3 ")
		n -= 5
	}
	a := [6]int{1, 2, -2, -1, 1, -1}
	for i := 1; i <= n; i++ {
		Fprint(out, a[(i-1)%6], " ")
	}
}

//func main() { cf1918G(os.Stdin, os.Stdout) }
