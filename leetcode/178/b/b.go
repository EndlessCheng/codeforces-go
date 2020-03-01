package main

import "sort"

func rankTeams(votes []string) (ans string) {
	bs := make([]byte, 26)
	for i := range bs {
		bs[i] = byte(i + 'A')
	}
	cnts := ['Z' + 1][26]int{}
	shown := ['Z' + 1]bool{}
	for _, s := range votes {
		for i, c := range s {
			cnts[c][i]++
			shown[c] = true
		}
	}
	sort.Slice(bs, func(i, j int) bool {
		a, b := bs[i], bs[j]
		if !shown[a] {
			return false
		}
		if !shown[b] {
			return true
		}
		ca, cb := cnts[a], cnts[b]
		for k := range ca {
			if ca[k] != cb[k] {
				return ca[k] > cb[k]
			}
		}
		return a < b
	})
	return string(bs[:len(votes[0])])
}
