package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1583F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)

	ans := 0
	for p := 1; p < n; p *= k {
		ans++
	}
	Fprintln(out, ans)

	for i := range n {
		for j := i + 1; j < n; j++ {
			c := 0
			for x, y := i, j; x != y; {
				c++
				x /= k
				y /= k
			}
			Fprint(out, c, " ")
		}
	}
}

//func main() { cf1583F(bufio.NewReader(os.Stdin), os.Stdout) }
