package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1854A2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		neg, pos, maxI, minI := 0, 0, 0, 0
		for i, v := range a {
			if v < 0 {
				neg++
			} else if v > 0 {
				pos++
			}
			if v > a[maxI] {
				maxI = i
			} else if v < a[minI] {
				minI = i
			}
		}

		type pair struct{ i, j int }
		ans := []pair{}
		makePos := func() {
			for i, v := range a {
				if v < 0 {
					ans = append(ans, pair{i, maxI})
				}
			}
			for i := 1; i < n; i++ {
				ans = append(ans, pair{i, i - 1})
			}
		}
		makeNeg := func() {
			for i, v := range a {
				if v > 0 {
					ans = append(ans, pair{i, minI})
				}
			}
			for i := n - 2; i >= 0; i-- {
				ans = append(ans, pair{i, i + 1})
			}
		}

		if a[maxI] >= -a[minI] {
			if neg <= 12 {
				makePos()
			} else {
				for i := 0; i < 5; i++ {
					ans = append(ans, pair{minI, minI})
				}
				makeNeg()
			}
		} else {
			if pos <= 12 {
				makeNeg()
			} else {
				for i := 0; i < 5; i++ {
					ans = append(ans, pair{maxI, maxI})
				}
				makePos()
			}
		}

		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p.i+1, p.j+1)
		}
	}
}

//func main() { cf1854A2(os.Stdin, os.Stdout) }
