package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2245B(in io.Reader, out io.Writer) {
	var T, n, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &c)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)

		i := min(sort.SearchInts(a, c), n/2)
		ans := -c * (n - i)
		for _, v := range a[i:] {
			ans += v
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2245B(bufio.NewReader(os.Stdin), os.Stdout) }
