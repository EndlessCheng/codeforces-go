package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1199B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var h, l float64
	Fscan(in, &h, &l)
	Fprintf(out, "%.13f", (l*l-h*h)/(2*h))
}

//func main() {
//	Sol1199B(os.Stdin, os.Stdout)
//}
