package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1774D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, n)
		cnt := make([]int, n)
		s := 0
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
				if a[i][j] > 0 {
					cnt[i]++
					s++
				}
			}
		}
		if s%n > 0 {
			Fprintln(out, -1)
			continue
		}

		ans := [][3]int{}
		avg := s / n
		for j := range m {
			var ins, outs []int
			for i, c := range cnt {
				if c < avg && a[i][j] == 0 {
					ins = append(ins, i)
				} else if c > avg && a[i][j] > 0 {
					outs = append(outs, i)
				}
			}
			for i := range min(len(ins), len(outs)) {
				ans = append(ans, [3]int{ins[i] + 1, outs[i] + 1, j + 1})
				cnt[ins[i]]++
				cnt[outs[i]]--
			}
		}

		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0], p[1], p[2])
		}
	}
}

//func main() { cf1774D(bufio.NewReader(os.Stdin), os.Stdout) }
