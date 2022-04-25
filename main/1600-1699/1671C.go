package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1671C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)

		ans, sum := int64(0), 0
		for i, v := range a {
			sum += v
			if sum > x {
				break
			}
			ans += int64((x-sum)/(i+1) + 1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1671C(os.Stdin, os.Stdout) }
