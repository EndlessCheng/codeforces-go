package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1714F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, d12, d23, d13 int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &d12, &d23, &d13)
		dCommon := (d12 + d13 - d23) / 2
		if d12+d23 < d13 || d12+d13 < d23 || d23+d13 < d12 || (d12+d23+d13)%2 > 0 || d12+d13-dCommon >= n {
			Fprintln(out, "NO")
			continue
		}
		type pair struct{ v, w int }
		ans := []pair{}
		cur := 3
		add := func(st, end, d int) {
			if d == 1 {
				ans = append(ans, pair{st, end})
			} else {
				cur++
				ans = append(ans, pair{st, cur})
				for i := 0; i < d-2; i++ {
					ans = append(ans, pair{cur, cur + 1})
					cur++
				}
				ans = append(ans, pair{cur, end})
			}
		}
		if dCommon == 0 {
			add(1, 2, d12)
			add(1, 3, d13)
		} else if dCommon < d12 && dCommon < d13 {
			cur = 4
			add(1, 4, dCommon)
			add(4, 2, d12-dCommon)
			add(4, 3, d13-dCommon)
		} else if dCommon == d12 {
			add(1, 2, dCommon)
			add(2, 3, d23)
		} else {
			add(1, 3, dCommon)
			add(3, 2, d23)
		}
		for cur < n {
			cur++
			ans = append(ans, pair{1, cur})
		}
		Fprintln(out, "YES")
		for _, e := range ans {
			Fprintln(out, e.v, e.w)
		}
	}
}

//func main() { CF1714F(os.Stdin, os.Stdout) }
