package main

func isSolvable(words []string, result string) bool {
	indexMap := func(s string) map[byte]int {
		mp := map[byte]int{}
		id := 0
		for i := range s {
			b := s[i]
			if _, ok := mp[b]; !ok {
				mp[b] = id
				id++
			}
		}
		return mp
	}
	allS := result
	for _, word := range words {
		allS += word
	}
	idMap := indexMap(allS)
	n := len(idMap)

	var canZero, used [10]bool
	var weights [10]int
	calcWeight := func(s string, sign int) {
		w := 1
		for i := len(s) - 1; i >= 0; i-- {
			weights[idMap[s[i]]] += w * sign
			w *= 10
		}
	}
	for i := range canZero {
		canZero[i] = true
	}
	for _, w := range words {
		canZero[idMap[w[0]]] = false
		calcWeight(w, 1)
	}
	canZero[idMap[result[0]]] = false
	calcWeight(result, -1)

	var f func(int, int) bool
	f = func(cur, sum int) bool {
		if cur == n {
			return sum == 0
		}
		for i := 0; i < 10; i++ {
			if i == 0 && !canZero[cur] || used[i] {
				continue
			}
			used[i] = true
			if f(cur+1, sum+i*weights[cur]) {
				return true
			}
			used[i] = false
		}
		return false
	}
	return f(0, 0)
}
