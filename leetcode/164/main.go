package main

import (
	. "fmt"
	"sort"
	"strings"
)

func minTimeToVisitAllPoints(points [][]int) int {
	n := len(points)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0
	for i := 1; i < n; i++ {
		pi := points[i]
		p0 := points[i-1]
		ans += max(abs(pi[0]-p0[0]), abs(pi[1]-p0[1]))
	}
	return ans
}

func countServers(grid [][]int) int {
	ans := 0
	r := len(grid)
	for i, gi := range grid {
	outer:
		for j, gij := range gi {
			if gij == 1 {
				for k, gik := range gi {
					if k != j && gik == 1 {
						ans++
						continue outer
					}
				}
				for k := 0; k < r; k++ {
					if k != i && grid[k][j] == 1 {
						ans++
						continue outer
					}
				}
			}
		}
	}
	return ans
}

type trieNode struct {
	childIdx [26]int
	dupCnt   int
	val      int
}

type trie struct {
	nodes []*trieNode
}

func newTrie() *trie {
	return &trie{
		nodes: []*trieNode{{}},
	}
}

func (t *trie) put(s string, val int) {
	o := t.nodes[0]
	for _, c := range s {
		c -= 'a'
		if o.childIdx[c] == 0 {
			o.childIdx[c] = len(t.nodes)
			t.nodes = append(t.nodes, &trieNode{})
		}
		o = t.nodes[o.childIdx[c]]
	}
	o.dupCnt++
	if o.dupCnt == 1 {
		o.val = val
	}
}

func (t *trie) minPrefix(p string) int {
	o := t.nodes[0]
	for _, c := range p {
		idx := o.childIdx[c-'a']
		if idx == 0 {
			return -1
		}
		o = t.nodes[idx]
	}
	for o.dupCnt == 0 {
		for i := 0; i < 26; i++ {
			if idx := o.childIdx[i]; idx > 0 {
				o = t.nodes[idx]
				break
			}
		}
	}
	return o.val
}

func suggestedProducts(products []string, searchWord string) (ans [][]string) {
	t := newTrie()
	sort.Strings(products)
	for i, p := range products {
		t.put(p, i)
	}
	for i := 1; i <= len(searchWord); i++ {
		sug := []string{}
		prefix := searchWord[:i]
		idx := t.minPrefix(prefix)
		if idx == -1 {
			ans = append(ans, sug)
			continue
		}
		for j := idx; j < len(products); j++ {
			p := products[j]
			if !strings.HasPrefix(p, prefix) {
				break
			}
			sug = append(sug, p)
			if len(sug) == 3 {
				break
			}
		}
		ans = append(ans, sug)
	}
	return
}

func numWays(steps int, arrLen int) int {
	const mx int = 505
	const mod int = 1e9 + 7

	dp := [mx][mx]int{}
	vis := [mx][mx]bool{}

	var f func(pos, left int) int
	f = func(pos, left int) int {
		if left == 0 {
			if pos == 0 {
				return 1
			}
			return 0
		}
		if pos < 0 || pos >= arrLen {
			return 0
		}
		if vis[pos][left] {
			return dp[pos][left]
		}

		ans := f(pos-1, left-1) + f(pos, left-1) + f(pos+1, left-1)
		dp[pos][left] = ans % mod
		vis[pos][left] = true
		return dp[pos][left]
	}
	return f(0, steps)
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

	//Println(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
	//Println(suggestedProducts([]string{"havana"}, "havana"))
	//Println(suggestedProducts([]string{"bags","baggage","banner","box","cloths"}, "bags"))
	//Println(suggestedProducts([]string{"havana"}, "tatiana"))
	Println(numWays(3, 2))
	Println(numWays(2, 4))
	Println(numWays(4, 2))
	Println(numWays(27, 7))
}
