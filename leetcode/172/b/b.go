package main

import "strings"

func printVertically(s string) (ans []string) {
	bs := [205][205]byte{}
	splits := strings.Split(s, " ")
	for i, sp := range splits {
		for j := range sp {
			bs[j][i] = sp[j]
		}
	}
	for i := range bs {
		ok := false
		for j := 204; j >= 0; j-- {
			if bs[i][j] != 0 {
				ok = true
				for k := j; k >= 0; k-- {
					if bs[i][k] == 0 {
						bs[i][k] = ' '
					}
				}
				ans = append(ans, string(bs[i][:j+1]))
				break
			}
		}
		if !ok {
			break
		}
	}
	return
}
