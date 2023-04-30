package main

// https://space.bilibili.com/206214
func score(a []int) (res int) {
	for i, x := range a {
		if i > 0 && a[i-1] == 10 || i > 1 && a[i-2] == 10 {
			x *= 2
		}
		res += x
	}
	return
}

func isWinner(player1, player2 []int) int {
	s1, s2 := score(player1), score(player2)
	if s1 > s2 {
		return 1
	}
	if s1 < s2 {
		return 2
	}
	return 0
}
