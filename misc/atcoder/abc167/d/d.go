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
	var n, k int
	Fscan(in, &n, &k)
	next := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &next[i])
	}

	beforeCycle := []int{}
	cycle := []int{}
	vis := make([]int8, n+1)
	for v := 1; vis[v] < 2; v = next[v] {
		if vis[v] == 1 {
			cycle = append(cycle, v)
		} else {
			beforeCycle = append(beforeCycle, v)
		}
		vis[v]++
	}
	if k < len(beforeCycle) {
		Fprint(_w, beforeCycle[k])
	} else {
		k -= len(beforeCycle)
		Fprint(_w, cycle[k%len(cycle)])
	}
}

func main() { run(os.Stdin, os.Stdout) }
