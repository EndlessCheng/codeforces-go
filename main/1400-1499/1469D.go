package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1469D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	key := map[int]bool{2: true, 4: true, 16: true, 256: true, 65536: true}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := [][2]int{}
		vs := []int{}
		for i := 2; i < n; i++ {
			if key[i] {
				vs = append(vs, i)
			} else {
				ans = append(ans, [2]int{i, n})
			}
		}
		vs = append(vs, n)
		for i := len(vs) - 1; i > 0; i-- {
			ans = append(ans, [2]int{vs[i], vs[i-1]}, [2]int{vs[i], vs[i-1]})
		}
		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0], p[1])
		}
	}
}

//func main() { CF1469D(os.Stdin, os.Stdout) }
