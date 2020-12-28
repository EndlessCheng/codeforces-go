package main

// github.com/EndlessCheng/codeforces-go
func subarrayBitwiseORs(a []int) (ans int) {
	has, b := map[int]bool{}, []int{}
	for _, v := range a {
		has[v] = true
		for i := range b {
			b[i] |= v
			has[b[i]] = true
		}
		b = append(b, v)
		j := 0
		for i := 1; i < len(b); i++ {
			if b[j] != b[i] {
				j++
				b[j] = b[i]
			}
		}
		b = b[:j+1]
	}
	return len(has)
}
