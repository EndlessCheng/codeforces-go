package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1998B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		for range n {
			Fscan(in, &v)
			Fprint(out, v%n+1, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1998B(bufio.NewReader(os.Stdin), os.Stdout) }
