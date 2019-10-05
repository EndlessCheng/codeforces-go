package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1077D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	cnt := make([]int, 2e5+1)
	for i := 0; i < n; i++ {
		var v int
		Fscan(in, &v)
		cnt[v]++
	}

	times := sort.Search(n/k+1, func(times int) bool {
		// return true if we can't cut out `times` times
		if times <= 1 {
			return false
		}
		left := k
		for _, c := range cnt {
			left -= c / times
			if left <= 0 {
				return false
			}
		}
		return true
	})
	times--
	for v, c := range cnt {
		for c /= times; c > 0; c-- {
			Fprint(out, v, " ")
			k--
			if k == 0 {
				return
			}
		}
	}
}

//func main() {
//	Sol1077D(os.Stdin, os.Stdout)
//}
