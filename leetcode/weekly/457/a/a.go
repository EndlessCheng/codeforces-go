package main

import (
	"slices"
	"unicode"
)

// https://space.bilibili.com/206214
var businessLineToCategory = map[string]int{
	"electronics": 0,
	"grocery":     1,
	"pharmacy":    2,
	"restaurant":  3,
}

// 检查字符串是否非空，只包含字母、数字和下划线
func isValid(s string) bool {
	for _, c := range s {
		if c != '_' && !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return s != ""
}

func validateCoupons(code []string, businessLine []string, isActive []bool) (ans []string) {
	groups := [4][]string{}
	for i, s := range code {
		category, ok := businessLineToCategory[businessLine[i]]
		if ok && isActive[i] && isValid(s) {
			groups[category] = append(groups[category], s) // 相同类别的优惠码分到同一组
		}
	}

	for _, g := range groups {
		slices.Sort(g) // 每一组内部排序
		ans = append(ans, g...)
	}
	return
}
