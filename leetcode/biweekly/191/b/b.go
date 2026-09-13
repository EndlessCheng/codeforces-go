package main

// https://space.bilibili.com/206214
func countSpecialIntegers(nums []int) (ans int) {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

next:
	for _, p := range pos {
		if len(p) < 3 {
			continue
		}
		d := p[1] - p[0]
		for i := 2; i < len(p); i++ {
			if p[i]-p[i-1] != d {
				continue next
			}
		}
		ans++
	}
	return
}
