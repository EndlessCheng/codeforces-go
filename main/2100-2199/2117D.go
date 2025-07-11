package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2117D(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		d := a[1] - a[0]
		for i := 2; i < n; i++ {
			if a[i]-a[i-1] != d {
				Fprintln(out, "NO")
				continue o
			}
		}
		v := a[0] - d
		if d < 0 {
			v = a[n-1] + d
		}
		if v >= 0 && v%(n+1) == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf2117D(bufio.NewReader(os.Stdin), os.Stdout) }
