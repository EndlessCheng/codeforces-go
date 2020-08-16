package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1257D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var st [][18]int
	stInit := func(a []int) {
		n := len(a)
		st = make([][18]int, n)
		for i := range st {
			st[i][0] = a[i]
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				st[i][j] = max(st[i][j-1], st[i+(1<<(j-1))][j-1])
			}
		}
	}
	stQuery := func(l, r int) int {
		k := uint(bits.Len(uint(r-l+1)) - 1)
		return max(st[l][k], st[r-(1<<k)+1][k])
	}
	type hero struct{ power, endurance int }

	var t int
	for Fscan(in, &t); t > 0; t-- {
		var n, m int
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		stInit(a)
		Fscan(in, &m)
		heroes := make([]hero, m)
		for i := range heroes {
			Fscan(in, &heroes[i].power, &heroes[i].endurance)
		}
		sort.Slice(heroes, func(i, j int) bool {
			return heroes[i].endurance > heroes[j].endurance || heroes[i].endurance == heroes[j].endurance && heroes[i].power > heroes[j].power
		})

		maxPowers := make([]int, n+1)
		maxPower := 0
		for i := 0; i < m; {
			hi := heroes[i]
			maxPower = max(maxPower, hi.power)
			j := i + 1
			for ; j < m && heroes[j].endurance == hi.endurance; j++ {
			}
			nextEndurance := 0
			if j < m {
				nextEndurance = heroes[j].endurance
			}
			for k := hi.endurance; k > nextEndurance; k-- {
				maxPowers[k] = maxPower
			}
			i = j
		}
		ans := 0
		for l := 0; l < n; ans++ {
			maxKills := sort.Search(n-l+1, func(kills int) bool {
				if kills == 0 {
					return !true
				}
				maxMonsterPower := stQuery(l, l+kills-1)
				return !(maxMonsterPower <= maxPowers[kills])
			}) - 1
			if maxKills == 0 {
				ans = -1
				break
			}
			l += maxKills
		}
		Fprintln(out, ans)
	}
}

//func main() {
//	Sol1257D(os.Stdin, os.Stdout)
//}
