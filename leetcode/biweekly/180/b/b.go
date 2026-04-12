package main

import (
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func countDigitOccurrences(nums []int, digit int) (ans int) {
	target := string('0' + byte(digit))
	for _, x := range nums {
		ans += strings.Count(strconv.Itoa(x), target)
	}
	return
}

func countDigitOccurrences2(nums []int, digit int) (ans int) {
	for _, x := range nums {
		for ; x > 0; x /= 10 {
			if x%10 == digit {
				ans++
			}
		}
	}
	return
}
