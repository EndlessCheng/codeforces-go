package main

/*一次遍历

由于每株植物都需要浇水，所以答案至少为植物的个数。我们只需要额外计算出在哪些位置往返即可，在位置 $i$ 处往返需要走 $2i$ 步，额外加上这些步数即为答案。
*/

// github.com/EndlessCheng/codeforces-go
func wateringPlants(plants []int, capacity int) int {
	ans := len(plants)
	water := capacity // 初始水量
	for i, need := range plants {
		if water < need {
			ans += 2 * i     // 往返
			water = capacity // 重置水量
		}
		water -= need
	}
	return ans
}
