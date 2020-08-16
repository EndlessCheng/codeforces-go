package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

func Sol1202D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var t, n, len3 int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		for len3 = 2; (len3+1)*len3 <= 2*n; len3++ {
		}
		n -= len3 * (len3 - 1) / 2
		Fprintln(out, "133"+strings.Repeat("7", n)+strings.Repeat("3", len3-2)+"7")
	}
}

//func main() {
//	Sol1202D(os.Stdin, os.Stdout)
//}
