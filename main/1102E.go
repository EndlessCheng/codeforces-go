package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Sol1102E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	arr := make([]int, n)
	rightI := make(map[int]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &arr[i])
		rightI[arr[i]] = i
	}

	ans := 1
	for l, r := 0, 0; r < n; {
		checkVal := arr[r]
		for _, val := range arr[l:rightI[checkVal]] {
			r = max(r, rightI[val])
		}
		if r+1 == n {
			break
		}
		ans = (ans << 1) % 998244353
		l = rightI[checkVal] + 1
		r++
	}
	Fprintln(out, ans)
}

func main() {
	Sol1102E(os.Stdin, os.Stdout)
}
