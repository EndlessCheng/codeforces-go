package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1067C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fprintln(out, 2*i/3, 2*i%3)
	}
}

//func main() { cf1067C(bufio.NewReader(os.Stdin), os.Stdout) }
