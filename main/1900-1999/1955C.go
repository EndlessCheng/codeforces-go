package main

import (
	. "fmt"
	"io"
)

func cf1955C(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		left := k
		for i := range a {
			Fscan(in, &a[i])
			left -= a[i]
		}
		if left >= 0 {
			Fprintln(out, n)
			continue
		}

		i := 0
		for left = (k + 1) / 2; left >= a[i]; i++ {
			left -= a[i]
		}
		j := n - 1
		for left = k / 2; left >= a[j]; j-- {
			left -= a[j]
		}
		Fprintln(out, i+n-1-j)
	}
}

//func main() { cf1955C(bufio.NewReader(os.Stdin), os.Stdout) }
