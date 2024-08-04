package main

// https://space.bilibili.com/206214
func minFlips(a [][]int) (ans int) {
	m, n := len(a), len(a[0])
	for i, row := range a[:m/2] {
		row2 := a[m-1-i]
		for j, x := range row[:n/2] {
			cnt1 := x + row[n-1-j] + row2[j] + row2[n-1-j]
			ans += min(cnt1, 4-cnt1) // 全为 1 或全为 0
		}
	}

	if m%2 > 0 && n%2 > 0 {
		// 正中间的数必须是 0
		ans += a[m/2][n/2]
	}

	diff, cnt1 := 0, 0
	if m%2 > 0 {
		// 统计正中间这一排
		row := a[m/2]
		for j, x := range row[:n/2] {
			if x != row[n-1-j] {
				diff++
			} else {
				cnt1 += x * 2
			}
		}
	}

	if n%2 > 0 {
		// 统计正中间这一列
		for i, row := range a[:m/2] {
			if row[n/2] != a[m-1-i][n/2] {
				diff++
			} else {
				cnt1 += row[n/2] * 2
			}
		}
	}

	if diff > 0 {
		ans += diff
	} else {
		ans += cnt1 % 4
	}
	return
}
