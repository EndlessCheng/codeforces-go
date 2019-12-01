package main

func numOfBurgers(a int, b int) []int {
	ans := []int{}
	if a < 2*b || (a-2*b)%2 == 1 {
		return ans
	}
	x := (a - 2*b) / 2
	y := b - x
	if y < 0 {
		return ans
	}
	return []int{x, y}
}
