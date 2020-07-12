package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, minI, maxI int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < a[minI] {
			minI = i
		} else if a[i] > a[maxI] {
			maxI = i
		}
	}
	ans := [][]int{}
	d := 1
	if a[minI] < 0 && a[maxI] > 0 {
		if a[maxI] > -a[minI] {
			for i, v := range a {
				if v < 0 {
					ans = append(ans, []int{maxI + 1, i + 1})
				}
			}
		} else {
			for i, v := range a {
				if v > 0 {
					ans = append(ans, []int{minI + 1, i + 1})
				}
			}
			d = 0
		}
	} else if a[maxI] <= 0 {
		d = 0
	}
	if d == 1 {
		for i := 1; i < n; i++ {
			ans = append(ans, []int{i, i + 1})
		}
	} else {
		for i := n; i > 1; i-- {
			ans = append(ans, []int{i, i - 1})
		}
	}
	Fprintln(out, len(ans))
	for _, p := range ans {
		Fprintln(out, p[0], p[1])
	}
}

func main() { run(os.Stdin, os.Stdout) }
