package main

// https://space.bilibili.com/206214
func minNumberOfHours(eng, exp int, energy, experience []int) (ans int) {
	for i, x := range energy {
		if eng <= x {
			ans += x - eng + 1
			eng = x + 1
		}
		eng -= x
		y := experience[i]
		if exp <= y {
			ans += y - exp + 1
			exp = y + 1
		}
		exp += y
	}
	return
}
