package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1707B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		i := 0
		for i < len(a)-1 {
			if i > 0 {
				i-- // 退一个 0
			}
			for j := i; j < len(a)-1; j++ {
				a[j] = a[j+1] - a[j]
			}
			a = a[:len(a)-1]
			slices.Sort(a[i:])
			if a[len(a)-1] == 0 {
				break
			}
			for a[i] == 0 {
				i++
			}
		}
		Fprintln(out, a[i])
	}
}

//func main() { cf1707B(bufio.NewReader(os.Stdin), os.Stdout) }
