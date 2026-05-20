package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2196A(in io.Reader, out io.Writer) {
	var T, p, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &p, &q)
		if p < q && p*3 >= q*2 {
			Fprintln(out, "Bob")
		} else {
			Fprintln(out, "Alice")
		}
	}
}

//func main() { cf2196A(bufio.NewReader(os.Stdin), os.Stdout) }
