package main

import (
	. "fmt"
	"io"
	. "strings"
)

// https://github.com/EndlessCheng
func cf2194D(in io.Reader, out io.Writer) {
	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		tot := 0
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
				tot += a[i][j]
			}
		}

		s := 0
		for i, row := range a {
			for j := m - 1; j >= 0; j-- {
				s += row[j]
				if s*2 >= tot {
					Fprintln(out, tot*tot/4)
					Fprintln(out, Repeat("D", i)+Repeat("R", j)+"D"+Repeat("R", m-j)+Repeat("D", n-1-i))
					continue o
				}
			}
		}
	}
}

//func main() { cf2194D(bufio.NewReader(os.Stdin), os.Stdout) }
