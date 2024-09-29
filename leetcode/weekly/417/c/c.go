package main

// https://space.bilibili.com/206214
func countOfSubstrings(word string, k int) (ans int64) {
	const vowelMask = 1065233
	var cntVowel1, cntVowel2 ['u' - 'a' + 1]int
	sizeVowel1, sizeVowel2 := 0, 0 // 元音种类数
	cntConsonant1, cntConsonant2 := 0, 0
	left1, left2 := 0, 0
	for _, b := range word {
		b -= 'a'
		if vowelMask>>b&1 > 0 {
			if cntVowel1[b] == 0 {
				sizeVowel1++
			}
			cntVowel1[b]++
			if cntVowel2[b] == 0 {
				sizeVowel2++
			}
			cntVowel2[b]++
		} else {
			cntConsonant1++
			cntConsonant2++
		}

		for sizeVowel1 == 5 && cntConsonant1 >= k {
			out := word[left1] - 'a'
			if vowelMask>>out&1 > 0 {
				cntVowel1[out]--
				if cntVowel1[out] == 0 {
					sizeVowel1--
				}
			} else {
				cntConsonant1--
			}
			left1++
		}

		for sizeVowel2 == 5 && cntConsonant2 > k {
			out := word[left2] - 'a'
			if vowelMask>>out&1 > 0 {
				cntVowel2[out]--
				if cntVowel2[out] == 0 {
					sizeVowel2--
				}
			} else {
				cntConsonant2--
			}
			left2++
		}

		ans += int64(left1 - left2)
	}
	return
}

func f(word string, k int) (ans int64) {
	const vowelMask = 1065233
	cnt1 := ['u' - 'a' + 1]int{}
	size1 := 0 // 元音种类数
	cnt2 := 0
	left := 0
	for _, b := range word {
		b -= 'a'
		if vowelMask>>b&1 > 0 {
			if cnt1[b] == 0 {
				size1++
			}
			cnt1[b]++
		} else {
			cnt2++
		}
		for size1 == 5 && cnt2 >= k {
			out := word[left] - 'a'
			if vowelMask>>out&1 > 0 {
				cnt1[out]--
				if cnt1[out] == 0 {
					size1--
				}
			} else {
				cnt2--
			}
			left++
		}
		ans += int64(left)
	}
	return
}

func countOfSubstrings2(word string, k int) int64 {
	return f(word, k) - f(word, k+1)
}
