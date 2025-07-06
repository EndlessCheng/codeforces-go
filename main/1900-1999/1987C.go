package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1987C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := n - 2; i >= 0; i-- {
			a[i] = max(a[i], a[i+1]+1)
		}
		Fprintln(out, a[0])
	}
}

//func main() { cf1987C(bufio.NewReader(os.Stdin), os.Stdout) }
