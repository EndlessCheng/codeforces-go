package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2111B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, a, b, c int
	f := [12]int{1, 1}
	for i := 2; i < 12; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		for range m {
			Fscan(in, &a, &b, &c)
			if f[n] <= min(a, b, c) && f[n+1] <= max(a, b, c) {
				Fprint(out, "1")
			} else {
				Fprint(out, "0")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf2111B(bufio.NewReader(os.Stdin), os.Stdout) }
