package main

func numSplits(s string) (ans int) {
	var cl, cr [26]int
	var l, r int
	for _, b := range s {
		b -= 'a'
		if cr[b] == 0 {
			r++
		}
		cr[b]++
	}
	for _, b := range s {
		b -= 'a'
		if cl[b] == 0 {
			l++
		}
		cl[b]++
		if cl[b] == cr[b] {
			r--
		}
		if l == r {
			ans++
		}
	}
	return
}
