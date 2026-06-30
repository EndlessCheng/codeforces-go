package main

import (
	"slices"
	"strconv"
)

func sequentialDigits(low, high int) (ans []int) {
	for d := 1; d <= 9; d++ {
		x := d
		for i := d; i <= 9 && x <= high; i++ {
			if x >= low {
				ans = append(ans, x)
			}
			// 把 i+1 加到 x 的末尾
			x = x*10 + i + 1
		}
	}
	slices.Sort(ans)
	return
}

func sequentialDigits2(low, high int) (ans []int) {
	const digits = "123456789"
	minLen := len(strconv.Itoa(low))
	maxLen := len(strconv.Itoa(high))
	// 枚举 digits 的子串（先枚举子串长度，再枚举子串位置）
	for length := minLen; length <= maxLen; length++ {
		for r := length; r <= len(digits); r++ {
			x, _ := strconv.Atoi(digits[r-length : r])
			if low <= x && x <= high {
				ans = append(ans, x)
			}
		}
	}
	return
}

func sequentialDigits3(low, high int) (ans []int) {
	x0 := 12 // 第一个窗口
	pow10 := 10
	for length := 2; x0 <= high; length++ {
		pow10 *= 10
		x := x0
		for i := length; i <= 9 && x <= high; i++ {
			if x >= low {
				ans = append(ans, x)
			}
			// 窗口向右滑动，i+1 进入窗口，i+1-length 离开窗口
			x = x*10 + i + 1 - (i+1-length)*pow10
		}
		x0 = x0*10 + length + 1
	}
	return
}
