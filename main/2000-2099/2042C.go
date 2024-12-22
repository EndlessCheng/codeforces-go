package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2042C(in io.Reader, out io.Writer) {
	var T, n, k int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		a := []int{}
		sum := 0
		for i := n - 1; i > 0; i-- {
			sum += int(s[i]-'0')*2 - 1
			if sum > 0 {
				a = append(a, sum)
			}
		}
		slices.Sort(a)

		ans := 1
		for i := len(a) - 1; i >= 0 && k > 0; i-- {
			k -= a[i]
			ans++
		}

		if k > 0 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, ans)
		}
	}
}

//func main() { cf2042C(bufio.NewReader(os.Stdin), os.Stdout) }
