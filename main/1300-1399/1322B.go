package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1322B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans, all int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		all ^= i
	}

	for k := 0; k < 25; k++ {
		mask := 1<<(k+1) - 1
		sort.Slice(a, func(i, j int) bool { return a[i]&mask < a[j]&mask })
		f := func(high int) (cnt int) {
			i, j := 0, n-1
			for i < j {
				if a[i]&mask+a[j]&mask < high {
					cnt ^= i ^ j
					i++
				} else {
					j--
				}
			}
			return
		}
		ans |= (f(1<<k) ^ f(1<<(k+1)) ^ f(3<<k) ^ all) & 1 << k
	}
	Fprint(out, ans)
}

//func main() { CF1322B(os.Stdin, os.Stdout) }
