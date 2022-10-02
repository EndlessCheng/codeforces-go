package main

// https://space.bilibili.com/206214
type LUPrefix struct {
	x   int
	has map[int]bool
}

func Constructor(int) LUPrefix {
	return LUPrefix{1, map[int]bool{}}
}

func (p LUPrefix) Upload(video int) {
	p.has[video] = true
}

func (p *LUPrefix) Longest() int {
	for p.has[p.x] {
		p.x++
	}
	return p.x - 1
}
