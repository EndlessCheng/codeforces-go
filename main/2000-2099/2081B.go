package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2081B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		cnt := 0
		l, r := -1, 0
		for i := range n - 1 {
			if a[i] > a[i+1] {
				cnt++
				if l == -1 {
					l = i
				}
				r = i + 1
			}
		}

		ans := cnt / 2
		if cnt%2 > 0 || l != -1 && r-l > a[r]-a[l] {
			ans++
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2081B(bufio.NewReader(os.Stdin), os.Stdout) }
