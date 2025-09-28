package main

// https://space.bilibili.com/206214
func majorityFrequencyGroup(s string) string {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	groups := map[int][]byte{}
	mx := 0
	for i, c := range cnt {
		if c == 0 {
			continue
		}
		groups[c] = append(groups[c], 'a'+byte(i))
		if len(groups[c]) > len(groups[mx]) || len(groups[c]) == len(groups[mx]) && c > mx {
			mx = c
		}
	}

	return string(groups[mx])
}
