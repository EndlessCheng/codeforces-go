package main

import (
	. "fmt"
	"io"
)

func cf1903C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans, s := 0, 0
		for i := n - 1; i >= 0; i-- {
			s += a[i]
			if i == 0 || s > 0 {
				ans += s
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1903C(bufio.NewReader(os.Stdin), os.Stdout) }
