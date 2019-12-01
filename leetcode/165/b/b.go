package main

func numOfBurgers(a int, b int) (ans []int) {
	ans = []int{}
	if a < 2*b || (a-2*b)%2 == 1 {
		return
	}
	x := (a - 2*b) / 2
	y := b - x
	if y < 0 {
		return
	}
	return []int{x, y}
}
