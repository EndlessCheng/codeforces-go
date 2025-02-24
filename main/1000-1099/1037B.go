package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1037B(in io.Reader, out io.Writer) {
	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	m := n / 2
	if a[m] > k {
		for i := m; i >= 0 && a[i] > k; i-- {
			ans += a[i] - k
		}
	} else {
		for i := m; i < n && a[i] < k; i++ {
			ans += k - a[i]
		}
	}
	Fprint(out, ans)
}

//func main() { cf1037B(bufio.NewReader(os.Stdin), os.Stdout) }
