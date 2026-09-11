package main

// https://space.bilibili.com/206214
func totalNumbers1(digits []int) int {
	set := map[int]struct{}{}
	for i, a := range digits {
		if a%2 > 0 {
			continue
		}
		for j, b := range digits {
			if j == i {
				continue
			}
			for k, c := range digits {
				if c == 0 || k == i || k == j {
					continue
				}
				set[c*100+b*10+a] = struct{}{}
			}
		}
	}
	return len(set)
}

func totalNumbers(digits []int) (ans int) {
	cnt := [10]int{}
	for _, d := range digits {
		cnt[d]++
	}

	var nonZeros, kinds, singles int
	for d, c := range cnt {
		if c == 0 {
			continue
		}
		kinds++
		if d > 0 {
			nonZeros++
			if c == 1 {
				singles++
			}
		}
	}

	// 枚举个位填偶数 d
	for d := 0; d < 10; d += 2 {
		c := cnt[d]
		if c == 0 {
			continue
		}

		// 十位填任意数字
		k := kinds
		if c == 1 {
			k--
		}

		// 百位填非零数字
		nz := nonZeros
		if d > 0 && c == 1 {
			nz--
		}

		// 恰好出现一次的非零数字，不能同时填入十位和百位
		s := singles
		if d > 0 {
			if c == 1 {
				s--
			} else if c == 2 {
				s++ // 个位数填入 d 后，d 恰好出现一次
			}
		}

		ans += k*nz - s
	}

	return
}
