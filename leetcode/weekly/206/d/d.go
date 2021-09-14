package main

// github.com/EndlessCheng/codeforces-go
func isTransformable(s, t string) (ans bool) {
	less := [10][][10]int{}
	cnt := [10]int{}
	for i := range s {
		b := s[i] & 15
		less[b] = append(less[b], cnt)
		cnt[b]++
	}

	cnt = [10]int{}
	for i := range t {
		b := t[i] & 15
		if len(less[b]) == 0 {
			return
		}
		for j, c := range less[b][0][:b] {
			if c > cnt[j] {
				return
			}
		}
		less[b] = less[b][1:]
		cnt[b]++
	}
	return true
}
