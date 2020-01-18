package main

import (
	"fmt"
)

// n, m := len(mat), len(mat[0])

func isSolvable(words []string, result string) bool {
	for i := range words {
		if len(words[i]) > len(result) {
			return false
		}
	}

	cs := make([][]byte, len(result))
	rs := make([]byte, len(result))
	for i := 0; i < len(result); i++ {
		for _, w := range words {
			j := len(w) - i - 1
			if j >= 0 {
				cs[i] = append(cs[i], w[j])
			}
		}
		rs[len(result)-i-1] = result[i]
	}

	ds := map[byte]int{}
	used := make([]bool, 10)
	var search func(int, int, int) bool
	search = func(i, j, carry int) bool {
		if i == len(result) {
			return carry == 0
		}
		if j == len(cs[i]) {
			s := carry
			for _, c := range cs[i] {
				s += ds[c]
			}
			s, carry = s%10, s/10
			if d, ok := ds[rs[i]]; ok {
				if d != s {
					return false
				}
				return search(i+1, 0, carry)
			} else if !used[s] {
				used[s] = true
				ds[rs[i]] = s
				defer func() {
					used[s] = false
					delete(ds, rs[i])
				}()
				return search(i+1, 0, carry)
			} else {
				return false
			}
		}
		c := cs[i][j]
		if _, ok := ds[c]; ok {
			return search(i, j+1, carry)
		}
		for d := 0; d < 10; d++ {
			if used[d] {
				continue
			}
			used[d] = true
			ds[c] = d
			res := search(i, j+1, carry)
			used[d] = false
			delete(ds, c)
			if res {
				return true
			}
		}
		return false
	}
	return search(0, 0, 0)
}

func longestCommonSubsequence(s1, s2 string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const mx = 1005
	dp := [mx][mx]int{}
	vis := [mx][mx]bool{}
	var f func(l, r int) int
	f = func(i, j int) (ans int) {
		if i == len(s1) || j == len(s2) {
			return 0
		}
		if vis[i][j] {
			return dp[i][j]
		}
		vis[i][j] = true
		defer func() { dp[i][j] = ans }()
		if s1[i] == s2[j] {
			return f(i+1, j+1) + 1
		}
		f1 := f(i+1, j)
		f2 := f(i, j+1)
		return max(f1, f2)
	}
	return f(0, 0)
}

func canVillagersWin(_players []string, credibility []int) bool {
	type player struct {
		id int
		name string
		c int
	}
	players := make([]player, len(_players))
	bearSeen := func() bool {
		for _, p := range players {
			if p.name == "bear" {
				return p.c == 100
			}
		}
		return false
	}

	for i, p := range _players {
		players[i] = player{
			id:   i,
			name: p,
			c:    credibility[i],
		}
	}

	// 1
	var killedPlayerIdx int
	if bearSeen() {
		for i, p := range players {
			if p.name == "bear" {
				killedPlayerIdx = i
				break
			}
		}
	} else {
		maxC := 0
		for i, p := range players {
			if p.name != "ww" && p.c > maxC {
				maxC = p.c
				killedPlayerIdx = i
			}
		}
	}
	players = append(players[:killedPlayerIdx], players[killedPlayerIdx+1:]...)

	// 2

	return false
}

func main() {
	toBytes := func(g [][]string) [][]byte {
		n, m := len(g), len(g[0])
		bytes := make([][]byte, n)
		for i := range bytes {
			bytes[i] = make([]byte, m)
			for j := range bytes[i] {
				bytes[i][j] = g[i][j][0]
			}
		}
		return bytes
	}
	_ = toBytes

	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func collections() {
	const mod int = 1e9 + 7
	dir4 := [...][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	ifElseI := func(cond bool, r1, r2 int) int {
		if cond {
			return r1
		}
		return r2
	}
	ifElseS := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}

	_ = []interface{}{fmt.Print, ifElseI, ifElseS, dir4, min, max, abs}
}
