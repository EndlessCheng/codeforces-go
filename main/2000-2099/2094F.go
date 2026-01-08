package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2094F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		if m%k > 0 {
			v := 0
			for range n {
				for range m {
					Fprint(out, v%k+1, " ")
					v++
				}
				Fprintln(out)
			}
		} else {
			for i := range n {
				for j := range m {
					Fprint(out, (i+j)%k+1, " ")
				}
				Fprintln(out)
			}
		}
	}
}

//func main() { cf2094F(bufio.NewReader(os.Stdin), os.Stdout) }
