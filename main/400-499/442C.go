package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf442C(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n+1)
	Fscan(in, &a[1], &a[2])
	for i := 3; i <= n; i++ {
		Fscan(in, &a[i])
		for a[i-2] >= a[i-1] && a[i-1] <= a[i] {
			ans += min(a[i-2], a[i])
			a[i-1] = a[i]
			i--
			n--
		}
	}
	slices.Sort(a[1 : n+1])
	for _, v := range a[1 : n-1] {
		ans += v
	}
	Fprint(out, ans)
}

//func main() { cf442C(bufio.NewReader(os.Stdin), os.Stdout) }
