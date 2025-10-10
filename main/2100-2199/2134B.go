package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2134B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		for range n {
			Fscan(in, &v)
			Fprint(out, v+v%(k+1)*k, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2134B(bufio.NewReader(os.Stdin), os.Stdout) }
