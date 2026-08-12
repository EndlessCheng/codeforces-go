package main

// github.com/EndlessCheng/codeforces-go
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	seats := map[int]int{} // 2~9 有预定座位的行 -> 这一行具体哪些座位被预定
	for _, r := range reservedSeats {
		seat := r[1]
		if 2 <= seat && seat <= 9 {
			seats[r[0]] |= 1 << (seat - 2) // 把二进制数的 seat-2 这一位变成 1
		}
	}

	// 注意：如果某一行只有 1 和 10 被预定，那么这一行不会插到哈希表中（相当于这一行是空的）
	// 示例 1 只有第 1 行和第 2 行插到哈希表中
	emptyRows := n - len(seats)
	ans := emptyRows * 2 // 一个空行可以容量 2 个四人小组
	for _, row := range seats {
		// 在哈希表中的行，由于 2~9 至少一个座位被预定，所以至多容纳 1 个四人小组，ans 至多增加 1
		if row&0b1111 == 0 || row&0b111100 == 0 || row&0b11110000 == 0 {
			ans++
		}
	}
	return ans
}
