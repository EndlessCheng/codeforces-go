package main

//func isSolvable(wordsS []string, resultS string) bool {
//	maxLen := 0
//	for _, w := range wordsS {
//		if len(w) > maxLen {
//			maxLen = len(w)
//		}
//	}
//	if maxLen > len(resultS) {
//		return false
//	}
//
//	wordsS = append(wordsS, resultS)
//
//	cnt := byte(0)
//	mp := map[byte]byte{}
//
//	words := make([][]byte, len(wordsS))
//	for i, w := range wordsS {
//		words[i] = make([]byte, 7)
//		for j := range words[i] {
//			words[i][j] = 99
//		}
//		st := 7 - len(w)
//		for j := st; j < 7; j++ {
//			words[i][j] = w[j-st] - 'A'
//		}
//	}
//
//	for j := 6; j >= 0; j-- {
//		for i := range words {
//			if words[i][j] != 99 {
//				if _, ok := mp[words[i][j]]; !ok {
//					mp[words[i][j]] = cnt
//					cnt++
//				}
//			}
//		}
//	}
//	for _, w := range words {
//		for j := range w {
//			if w[j] != 99 {
//				w[j] = mp[w[j]]
//			}
//		}
//	}
//	for i := range words {
//		s := string(words[i])
//		tmp := []byte(s)
//		words[i] = make([]byte, 7)
//		for j := range words[i] {
//			words[i][j] = 99
//		}
//		copy(words[i][7-len(tmp):], tmp)
//	}
//
//	kinds := len(mp)
//	notZero := make([]bool, kinds)
//	for _, s := range words {
//		notZero[s[0]] = true
//	}
//
//	cols := [7][]byte{}
//	for i := range cols {
//		for _, c := range words[i] {
//			if c != 99 {
//				cols[i] = append(cols[i], c)
//			}
//		}
//	}
//	maxCol := [7]int{}
//	for j :=range cols[0] {
//		if len(cols[j])
//		for i := range cols {
//
//		}
//	}
//
//	checkCol := func(col []byte, val []int, st int) bool {
//		col, res := col[:len(col)-1], col[len(col)-1]
//		sum := st
//		for _, c := range col {
//			sum += val[c]
//		}
//		return sum%10 == val[res]
//	}
//
//	var f func(dep int) bool
//	f = func(dep int) bool {
//		if dep == kinds {
//			checkCol()
//		}
//
//		for i := 0; i < 10; i++ {
//
//		}
//
//	}
//
//	return f(0)
//}
