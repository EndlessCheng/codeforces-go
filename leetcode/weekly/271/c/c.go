package main

// 双指针模拟

// github.com/EndlessCheng/codeforces-go
func minimumRefill(plants []int, capacityA, capacityB int) (ans int) {
	a, b := capacityA, capacityB
	for i, j := 0, len(plants)-1; i <= j; {
		if i == j {
			if a < plants[i] && b < plants[i] { // 两人都无法浇水
				ans++
			}
			break
		}
		if a < plants[i] {
			a = capacityA
			ans++
		}
		a -= plants[i]
		i++
		if b < plants[j] {
			b = capacityB
			ans++
		}
		b -= plants[j]
		j--
	}
	return
}
