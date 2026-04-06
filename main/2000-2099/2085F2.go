package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2085F2(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		last := make([]int, n+1)
		diff := make([]int, n+1)
		s := 0
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			diff[i] += 2
			if last[v] == 0 {
				s += i - 1
			} else {
				diff[(i+last[v])/2]--
				diff[(i+last[v]+1)/2]--
			}
			last[v] = i
		}

		ans := int(1e18)
		sd := -k
		for i := 1; i <= n; i++ {
			sd += diff[i]
			s += sd
			ans = min(ans, s)
		}
		for i := 1; i <= k; i++ {
			ans -= abs(i - (k+1)/2)
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2085F2(bufio.NewReader(os.Stdin), os.Stdout) }
