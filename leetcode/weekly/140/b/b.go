package main

func numTilePossibilities(tiles string) int {
	calcP := func(n, k int) int {
		res := 1
		for i := n; i > n-k; i-- {
			res *= i
		}
		return res
	}

	cnt := [26]int{}
	for _, c := range tiles {
		cnt[c-'A']++
	}
	ans := 0
	for _, c := range cnt {
		if c > 0 {
			ans++
		}
	}
	n := len(tiles)
	for i := 2; i <= n; i++ {
		// A(n,i)/(PI(A(c,c)))
		res := calcP(n, i)
		for _, c := range cnt {
			if c > 1 {
				res /= calcP(c, c)
			}
		}
		ans += res
	}
	return ans
}
