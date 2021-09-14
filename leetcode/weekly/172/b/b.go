package main

import "strings"

func printVertically(s string) (ans []string) {
	bs := [205][205]byte{}
	for i := range bs {
		for j := range bs[i] {
			bs[i][j] = ' '
		}
	}
	splits := strings.Split(s, " ")
	for i, sp := range splits {
		for j := range sp {
			bs[j][i] = sp[j]
		}
	}
	for i := range bs {
		for j := 204; j >= 0; j-- {
			if bs[i][j] != ' ' {
				ans = append(ans, string(bs[i][:j+1]))
				break
			}
		}
	}
	return
}
