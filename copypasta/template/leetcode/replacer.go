package leetcode

import "strings"

var inputNameReplacer *strings.Replacer

func init() {
	// 替换常见变量名（数组、字符串等）
	oldNew := []string{
		// 数组、矩阵
		"nums", "a",
		"nums1", "a",
		"nums2", "b",
		"nums3", "c",
		"arr", "a",
		"brr", "b",
		"array", "a",
		"brray", "b",
		"elements", "a",
		"scores", "a",
		"values", "a",
		"vals", "a",
		"stones", "a",
		"cards", "a",
		"prices", "a",
		"sales", "a",
		"grades", "a",
		"beans", "a",
		"coins", "a",
		"mat", "a",
		"matrix", "a",
		"grid", "a",
		"grid1", "a",
		"grid2", "b",
		"words", "a",
		"events", "a",
		"heights", "a",

		// 字符串
		"word", "s",
		"word1", "s",
		"word2", "t",
		"word3", "z",
		"s1", "s",
		"s2", "t",
		"s3", "z",

		// 其余常见变量名
		"number", "n",
		"num", "n",
		"num1", "x",
		"num2", "y",
		"num3", "z",
		"size", "n",
		"edges", "es",
		"edges1", "es1",
		"edges2", "es2",
		"points", "ps",
		"point1", "p1",
		"point2", "p2",
		"point3", "p3",
		"pairs", "ps",
		"query", "qs",
		"queries", "qs",
		"startPos", "st",
		"start", "st",
		"source", "st",
		"destination", "end",
		"target", "tar",
		"total", "tot",
		"limit", "lim",
		"index", "id",
		"index1", "id1",
		"index2", "id2",
		"dist", "dis",
		"timestamp", "ts",
		"diff", "d",
		"event1", "e1",
		"event2", "e2",
		"lower", "low",
	}
	for i := range oldNew {
		oldNew[i] += " " // 由于要匹配变量名+空格+类型，为了防止修改到意外的位置，通过加一个空格来简单地实现匹配
	}
	inputNameReplacer = strings.NewReplacer(oldNew...)
	inputNameReplacer.Replace("") // 触发内部的 buildOnce
}

func renameInputArgs(funcDefineLine string) string {
	return inputNameReplacer.Replace(funcDefineLine)
}
