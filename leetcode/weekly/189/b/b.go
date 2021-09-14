package main

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	sp := strings.Split(text, " ")
	sp[0] = strings.ToLower(sp[0])
	sort.SliceStable(sp, func(i, j int) bool { return len(sp[i]) < len(sp[j]) })
	sp[0] = strings.Title(sp[0])
	return strings.Join(sp, " ")
}
