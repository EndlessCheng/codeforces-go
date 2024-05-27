package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1066B(in io.Reader, out io.Writer) {
	var n, r, v, ans int
	Fscan(in, &n, &r)
	a := []int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v > 0 {
			a = append(a, i)
		}
	}
	done := 0
	for done < n {
		i := sort.SearchInts(a, done+r) - 1
		if i < 0 {
			Fprint(out, -1)
			return
		}
		ans++
		done = a[i] + r
		a = a[i+1:]
	}
	Fprint(out, ans)
}

//func main() { cf1066B(bufio.NewReader(os.Stdin), os.Stdout) }
