package copypasta

import "testing"

func Test_graph(t *testing.T) {
	n := 10
	g := newGraph(n)
	g.addBoth(1, 2, 1)
	g.addBoth(2, 3, 1)
	g.addBoth(3, 4, 1)
	g.addBoth(3, 5, 1)
	g.addBoth(5, 6, 1)

	calc := func(start int) (anotherStart int, maxPath int) {
		const inf = 1e9
		dis := make([]int, n+1)
		for i := range dis {
			dis[i] = inf
		}
		dis[start] = 0
		g.reset()
		g.dfs(start, func(from, to int, weight int) {
			dis[to] = dis[from] + weight
		})
		for v := range dis {
			if dis[v] != inf && dis[v] > maxPath {
				maxPath = dis[v]
				anotherStart = v
			}
		}
		return
	}
	s0 := 3
	s1, _ := calc(s0)
	s2, ans := calc(s1)
	t.Log(s0, s1, s2, ans)
}
