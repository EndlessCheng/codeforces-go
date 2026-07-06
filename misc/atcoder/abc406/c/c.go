package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n+1)
	pos := []int{0}
	for i := range a {
		Fscan(in, &a[i])
		if i > 1 && (a[i-2] < a[i-1]) == (a[i-1] > a[i]) {
			pos = append(pos, i-1)
			m := len(pos)
			if m > 3 && a[i-1] > a[i] {
				ans += (pos[m-3] - pos[m-4]) * (i - 1 - pos[m-2])
			}
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
