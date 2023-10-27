package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF289B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	
	var n, m, x, ans int
	Fscan(in, &n, &m, &x)
	a := make([]int, n*m)
	for i := range a {
		Fscan(in, &a[i])
		if (a[i]-a[0])%x != 0 {
			Fprint(out, -1)
			return
		}
	}
	sort.Ints(a)
	for _, v := range a {
		ans += abs(v-a[n*m/2]) / x
	}
	Fprint(out, ans)
}

//func main() { CF289B(os.Stdin, os.Stdout) }
