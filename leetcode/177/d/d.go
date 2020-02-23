package main

func largestMultipleOfThree(digits []int) string {
	cnts := [10]int{}
	sum := 0
	for _, d := range digits {
		cnts[d]++
		sum += d
	}

	func() {
		if sum%3 == 0 {
			return
		}

		// 另一种方法是暴力 for 判断
		check1, check2 := [3]int{1, 4, 7}, [6][2]int{{2, 2}, {2, 5}, {5, 5}, {2, 8}, {5, 8}, {8, 8}}
		if sum%3 == 2 {
			check1, check2 = [3]int{2, 5, 8}, [6][2]int{{1, 1}, {1, 4}, {4, 4}, {1, 7}, {4, 7}, {7, 7}}
		}
		for _, v := range check1 {
			if cnts[v] > 0 {
				cnts[v]--
				return
			}
		}
		for _, p := range check2 {
			if a, b := p[0], p[1]; a == b {
				if cnts[a] > 1 {
					cnts[a] -= 2
					return
				}
			} else if cnts[a] > 0 && cnts[b] > 0 {
				cnts[a]--
				cnts[b]--
				return
			}
		}
	}()

	digits = []int{}
	for d, c := range cnts {
		for ; c > 0; c-- {
			digits = append(digits, d)
		}
	}
	n := len(digits)
	if n == 0 {
		return ""
	}
	if cnts[0] == n {
		return "0"
	}
	ans := []byte{}
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, byte('0'+digits[i]))
	}
	return string(ans)
}
