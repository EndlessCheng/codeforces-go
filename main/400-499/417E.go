package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf417E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	f := func(i, x int) int {
		if i < x {
			return 4
		}
		if x < 3 {
			return 3
		}
		return (x - 2) << 1
	}

	var n, m int
	Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			Fprint(out, f(i, n)*f(j, m), " ")
		}
		Fprintln(out)
	}
}

//func main() { cf417E(bufio.NewReader(os.Stdin), os.Stdout) }
