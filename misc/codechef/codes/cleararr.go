package main

import (
	. "fmt"
	"io"
	"sort"
)

func clearArr(in io.Reader, out io.Writer) {
	var T, n, k, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &x)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0
		sort.Ints(a)
		i := n - 1
		for ; k > 0 && i > 0 && a[i]+a[i-1] > x; i -= 2 {
			ans += x
			k--
		}
		for ; i >= 0; i-- {
			ans += a[i]
		}
		Fprintln(out, ans)
	}
}

//func main() { clearArr(bufio.NewReader(os.Stdin), os.Stdout) }
