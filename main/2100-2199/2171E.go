package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2171E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)

		a := []int{}
		for x := 1; x <= n; x += 6 {
			a = append(a, x)
		}
		for x := 5; x <= n; x += 6 {
			a = append(a, x)
		}

		b := []int{}
		for x := 2; x <= n; x += 2 {
			b = append(b, x)
		}
		for x := 3; x <= n; x += 6 {
			b = append(b, x)
		}

		i, j := 0, 0
		for i < len(a) && j+1 < len(b) {
			Fprint(out, a[i], b[j], b[j+1], " ")
			i++
			j += 2
		}
		for _, v := range append(a[i:], b[j:]...) {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2171E(bufio.NewReader(os.Stdin), os.Stdout) }
