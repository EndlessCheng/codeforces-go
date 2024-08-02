package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf1949B(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		slices.Sort(b)
		ans := 0
		for k := 0; k < n; k++ {
			mn := int(1e9)
			for i, v := range a {
				mn = min(mn, abs(v-b[(i+k)%n]))
			}
			ans = max(ans, mn)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1949B(bufio.NewReader(os.Stdin), os.Stdout) }
