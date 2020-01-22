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

	calcWeight := func(s string, a *[10]int) {
		w := 1
		for i := len(s) - 1; i >= 0; i-- {
			a[idMap[s[i]]] += w
			w *= 10
		}
	}
	var canZero, used [10]bool
	var lWeights, rWeights [10]int
	for i := range canZero {
		canZero[i] = true
	}
	for _, w := range words {
		canZero[idMap[w[0]]] = false
		calcWeight(w, &lWeights)
	}
	canZero[idMap[result[0]]] = false
	calcWeight(result, &rWeights)

	var f func(cur, l, r int) bool
	f = func(cur, l, r int) bool {
		if cur == n {
			return l == r
		}
		for i := 0; i < 10; i++ {
			if i == 0 && !canZero[cur] || used[i] {
				continue
			}
			used[i] = true
			if f(cur+1, l+i*lWeights[cur], r+i*rWeights[cur]) {
				return true
			}
			used[i] = false
		}
		return false
	}
	return f(0, 0, 0)
}
