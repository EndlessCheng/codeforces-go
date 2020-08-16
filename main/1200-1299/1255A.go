package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1255A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var t, a, b int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b)
		d := b - a
		if d < 0 {
			d = -d
		}
		Fprintln(out, d/5+(d%5+1)/2)
	}
}

//func main() {
//	Sol1255A(os.Stdin, os.Stdout)
//}
