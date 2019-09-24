package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol1196E(reader io.Reader, writer io.Writer) {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	for Fscan(in, &n); n > 0; n-- {
		var b, w int
		Fscan(in, &b, &w)

		start := 2
		if b > w {
			start = 1
			b, w = w, b
		}
		if w > 3*b+1 {
			Fprintln(out, "NO")
			continue
		}

		Fprintln(out, "YES")
		for i := 0; i < b; i++ {
			Fprintln(out, 2, start+2*i)
			Fprintln(out, 2, start+2*i+1)
		}
		w -= b
		if w > 0 {
			Fprintln(out, 2, start+2*b)
			w--
		}
		minC := min(w, b)
		for i := 0; i < minC; i++ {
			Fprintln(out, 1, start+2*i+1)
		}
		w -= minC
		for i := 0; i < w; i++ {
			Fprintln(out, 3, start+2*i+1)
		}
	}
}

//func main() {
//	Sol1196E(os.Stdin, os.Stdout)
//}
