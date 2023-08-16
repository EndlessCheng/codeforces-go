package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	m := n / 2
	ans := m * (m + 1) * (m*4 + n%2*6 - 1) / 6
	pos := make([][]int, n+1)
	sumP := make([][]int, n+1)
	for i := range sumP {
		sumP[i] = []int{0}
	}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		p := sort.SearchInts(pos[v], n-1-i)
		ans -= sumP[v][p] + (len(pos[v])-p)*(n-i)
		pos[v] = append(pos[v], i)
		sumP[v] = append(sumP[v], sumP[v][len(sumP[v])-1]+i+1)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
