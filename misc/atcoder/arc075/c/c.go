package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func mergeSort(a []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	left := append(a[:0:0], a[:n/2]...)
	right := append(a[:0:0], a[n/2:]...)
	cnt := mergeSort(left) + mergeSort(right)
	l, r := 0, 0
	for i := range a {
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			cnt += len(right) - r
			a[i] = left[l]
			l++
		} else {
			a[i] = right[r]
			r++
		}
	}
	return cnt
}

func run(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] += a[i-1] - k
	}
	Fprint(out, mergeSort(a))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
