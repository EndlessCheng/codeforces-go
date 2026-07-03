package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1208G(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = i
	}
	for i := 2; i <= n; i++ {
		if a[i] == i {
			for j := i; j <= n; j += i {
				a[j] = a[j] / i * (i - 1)
			}
		}
	}
	slices.Sort(a[3:])

	ans := 1
	if k > 1 {
		ans++
	}
	for _, v := range a[3 : k+3] {
		ans += v
	}
	Fprint(out, ans)
}

//func main() { cf1208G(bufio.NewReader(os.Stdin), os.Stdout) }
