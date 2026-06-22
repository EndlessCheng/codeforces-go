package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1421E(in io.Reader, out io.Writer) {
	var n, sum, ans int
	Fscan(in, &n)
	a := make([]int, n)
	s := [2]int{}
	for i := range a {
		Fscan(in, &a[i])
		sum += a[i]
		s[i%2] += a[i]
	}
	slices.Sort(a)

	ans = -9e18
	if n%3 == 1 {
		ans = sum
	}
	for i, v := range a {
		sum -= v * 2
		if (n+i)%3 == 0 {
			if sum == s[0]-s[1] || sum == s[1]-s[0] {
				ans = max(ans, sum+(a[i]-a[i+1])*2)
			} else {
				ans = max(ans, sum)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1421E(bufio.NewReader(os.Stdin), os.Stdout) }
