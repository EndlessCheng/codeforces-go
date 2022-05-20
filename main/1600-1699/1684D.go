package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1684D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ori := make([]int, n)
		type pair struct{ v, i int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &ori[i])
			a[i] = pair{ori[i] + i, i}
		}
		sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })
		del := make([]bool, n)
		for _, p := range a[:k] {
			del[p.i] = true
		}
		ans, ex := int64(0), 0
		for i, b := range del {
			if b {
				ex++
			} else {
				ans += int64(ori[i] + ex)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1684D(os.Stdin, os.Stdout) }
