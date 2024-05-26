package main

import "bytes"

// https://space.bilibili.com/206214
func compressedString(word string) string {
	t := []byte{}
	i0 := -1
	for i := range word {
		c := word[i]
		if i+1 == len(word) || c != word[i+1] {
			k := i - i0
			t = append(t, bytes.Repeat([]byte{'9', c}, k/9)...)
			if k%9 > 0 {
				t = append(t, '0'+byte(k%9), c)
			}
			i0 = i
		}
	}
	return string(t)
}
