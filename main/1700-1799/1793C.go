package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1793C(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		mn, mx := 1, n
		l, r := 0, n-1
		for l < r {
			if a[l] == mn {
				l++
				mn++
			} else if a[l] == mx {
				l++
				mx--
			} else if a[r] == mn {
				r--
				mn++
			} else if a[r] == mx {
				r--
				mx--
			} else {
				Fprintln(out, l+1, r+1)
				continue o
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { cf1793C(bufio.NewReader(os.Stdin), os.Stdout) }
