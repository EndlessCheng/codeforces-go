package main

import (
	. "fmt"
	"io"
	"maps"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2037F(in io.Reader, out io.Writer) {
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([]struct{ hp, x int }, n)
		for i := range a {
			Fscan(in, &a[i].hp)
		}
		for i := range a {
			Fscan(in, &a[i].x)
		}

		ans := 1 + sort.Search(1e9, func(atk int) bool {
			atk++
			diff := map[int]int{}
			for _, p := range a {
				d := m - 1 - (p.hp-1)/atk
				if d >= 0 {
					diff[p.x-d]++
					diff[p.x+d+1]--
				}
			}
			sumD := 0
			for _, x := range slices.Sorted(maps.Keys(diff)) {
				sumD += diff[x]
				if sumD >= k {
					return true
				}
			}
			return false
		})
		if ans > 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2037F(bufio.NewReader(os.Stdin), os.Stdout) }
