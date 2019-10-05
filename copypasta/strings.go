package copypasta

func stringCollection() {
	// code from my answer at https://www.zhihu.com/question/21923021/answer/37475572
	calcMaxMatchLengths := func(pattern string) []int {
		n := len(pattern)
		maxMatchLengths := make([]int, n)
		maxLength := 0
		for i := 1; i < n; i++ {
			c := pattern[i]
			for maxLength > 0 && pattern[maxLength] != c {
				maxLength = maxMatchLengths[maxLength-1]
			}
			if pattern[maxLength] == c {
				maxLength++
			}
			maxMatchLengths[i] = maxLength
		}
		return maxMatchLengths
	}
	// search pattern from text, return all start positions
	kmpSearch := func(text, pattern string) (positions []int) {
		maxMatchLengths := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		count := 0
		for i := range text {
			c := text[i]
			for count > 0 && pattern[count] != c {
				count = maxMatchLengths[count-1]
			}
			if pattern[count] == c {
				count++
			}
			if count == lenP {
				positions = append(positions, i-lenP+1)
				count = maxMatchLengths[count-1]
			}
		}
		return
	}

	calcMinPeriod := func(pattern string) int {
		maxMatchLengths := calcMaxMatchLengths(pattern)
		n := len(pattern)
		if val := maxMatchLengths[n-1]; val > 0 {
			if n%(n-val) == 0 {
				return n / (n - val)
			}
		}
		return 1 // or -1
	}

	_ = []interface{}{kmpSearch, calcMinPeriod}
}
