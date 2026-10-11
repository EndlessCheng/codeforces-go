package main

import "strconv"

// https://space.bilibili.com/206214
type BusBooking struct {
	states    []int
	times     [][4]int
	totalTime int
	bookedNum int
}

func Constructor(n int) BusBooking {
	return BusBooking{
		states: make([]int, n+1),
		times:  make([][4]int, n+1),
	}
}

func (b *BusBooking) Toggle(seat string) {
	row, _ := strconv.Atoi(seat[1:])
	mask := b.states[row]
	col := seat[0] - 'A'

	if mask>>col&1 == 0 {
		// 预订
		time := 1
		if col == 0 && mask>>1&1 > 0 || col == 3 && mask>>2&1 > 0 {
			time = 3 // 靠窗，且相邻座位已被预订
		}
		b.totalTime += time
		b.times[row][col] = time // 记录预订该座位所需的秒数，将来释放时用
		b.bookedNum++
	} else {
		// 释放
		b.totalTime -= b.times[row][col]
		b.times[row][col] = 0
		b.bookedNum--
	}

	b.states[row] ^= 1 << col
}

func (b *BusBooking) GetTotalTime() int {
	return b.totalTime
}

func (b *BusBooking) GetMinTime() int {
	return b.bookedNum
}
