package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://vjudge.net/contest/351746#problem/D
// github.com/EndlessCheng/codeforces-go
func Sol1077E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, ans int
	Fscan(in, &n)
	cntMp := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		cntMp[v]++
	}
	cnts := make([]int, 0, len(cntMp))
	for _, c := range cntMp {
		cnts = append(cnts, c)
	}
	sort.Ints(cnts)
	for st := 1; st <= n; st++ {
		i := 0
		for v = st; ; v <<= 1 {
			i += sort.SearchInts(cnts[i:], v) + 1
			if i > len(cnts) {
				break
			}
		}
		if sum := v - st; sum > ans {
			ans = sum
		}
	}
	Fprint(out, ans)
}

func main() {
	Sol1077E(os.Stdin, os.Stdout)
}
