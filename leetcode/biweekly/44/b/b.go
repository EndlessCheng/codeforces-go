package main

// github.com/EndlessCheng/codeforces-go
func minimumTeachings(n int, languages, friendships [][]int) (ans int) {
	ls := make([]map[int]bool, len(languages))
	for i, l := range languages {
		ls[i] = map[int]bool{}
		for _, v := range l {
			ls[i][v] = true
		}
	}

	ok := make([]bool, len(friendships))
	for i, e := range friendships {
		e[0]--
		e[1]--
		for v := range ls[e[0]] {
			if ls[e[1]][v] {
				ok[i] = true
				break
			}
		}
	}

	ans = 1e9
	for l := 1; l <= n; l++ {
		tech := map[int]bool{}
		for i, e := range friendships {
			if ok[i] {
				continue
			}
			v, w := e[0], e[1]
			if !ls[v][l] && !tech[v] {
				tech[v] = true
			}
			if !ls[w][l] && !tech[w] {
				tech[w] = true
			}
		}
		ans = min(ans, len(tech))
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
