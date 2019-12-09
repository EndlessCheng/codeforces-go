package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(cnts [][26]int, bits int) int {
	cnt := [26]bool{}

	for i := 0; bits > 0; bits >>= 1 {
		if bits&1 == 1 {

			for i, c := range cnts[i] {
				if c > 1 {
					return 0
				}
				if c > 0 {
					if !cnt[i] {
						cnt[i] = true
					} else {
						return 0
					}
				}
			}

		}
		i++
	}
	ans := 0
	for _, v := range cnt {
		if v {
			ans++
		}
	}
	return ans
}

func maxLength(arr []string) int {
	n := len(arr)
	cnts := make([][26]int, n)
	for i, s := range arr {
		for _, c := range s {
			cnts[i][c-'a']++
		}
	}

	ans := 0
	n2 := 1 << uint(n)
	for i := 1; i < n2; i++ {
		ans = max(ans, check(cnts, i))
	}
	return ans
}
