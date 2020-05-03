package main

func kLengthApart(a []int, k int) (ans bool) {
	pos := []int{}
	for i, v := range a {
		if v == 1 {
			pos = append(pos, i)
		}
	}
	for i := 1; i < len(pos); i++ {
		if pos[i]-pos[i-1]-1 < k {
			return false
		}
	}
	return true
}
