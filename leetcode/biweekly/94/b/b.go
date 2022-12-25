package main

import (
	"sort"
	"strings"
)

// https://space.bilibili.com/206214
func topStudents(positiveFeedback, negativeFeedback, report []string, studentId []int, k int) []int {
	score := map[string]int{}
	for _, w := range positiveFeedback {
		score[w] = 3
	}
	for _, w := range negativeFeedback {
		score[w] = -1
	}
	type pair struct{ score, id int }
	a := make([]pair, len(report))
	for i, r := range report {
		s := 0
		for _, w := range strings.Split(r, " ") {
			s += score[w]
		}
		a[i] = pair{s, studentId[i]}
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.score > b.score || a.score == b.score && a.id < b.id
	})
	ans := make([]int, k)
	for i, p := range a[:k] {
		ans[i] = p.id
	}
	return ans
}
