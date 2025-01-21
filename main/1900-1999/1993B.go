package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1993B(in io.Reader, out io.Writer) {
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := [2][]int{}
		for range n {
			Fscan(in, &v)
			a[v&1] = append(a[v&1], v)
		}
		if len(a[1]) == 0 {
			Fprintln(out, 0)
			continue
		}
		mx := slices.Max(a[1])
		slices.Sort(a[0])
		ans := len(a[0])
		for _, v := range a[0] {
			if v > mx {
				ans++
				break
			}
			mx += v
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1993B(bufio.NewReader(os.Stdin), os.Stdout) }
