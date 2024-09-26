package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2004F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		cnt := map[int]int{}
		for i := range a {
			s := 0
			for j := i; j < n; j++ {
				s += a[j]
				ans += j - i - cnt[s]
				cnt[s]++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2004F(bufio.NewReader(os.Stdin), os.Stdout) }
