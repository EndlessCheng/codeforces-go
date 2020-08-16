package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1255B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var t, n, m, a int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		sum := 0
		for i := 0; i < n; i++ {
			Fscan(in, &a)
			sum += a * 2
		}
		if n == 2 || n > m {
			Fprintln(out, -1)
			continue
		}
		Fprintln(out, sum)
		for i := 1; i < n; i++ {
			Fprintln(out, i, i+1)
		}
		Fprintln(out, n, 1)
	}
}

//func main() {
//	Sol1255B(os.Stdin, os.Stdout)
//}
