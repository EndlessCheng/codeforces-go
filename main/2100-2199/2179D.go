package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2179D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprint(out, 1<<n-1, " ")
		for i := n - 1; i >= 0; i-- {
			v := 1<<i - 1
			for j := range 1 << (n - 1 - i) {
				Fprint(out, j<<(i+1)|v, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2179D(os.Stdin, os.Stdout) }
