package main

// github.com/EndlessCheng/codeforces-go
type Encrypter struct {
	mp  [26]string
	cnt map[string]int
}

func Constructor(keys []byte, values, dictionary []string) Encrypter {
	mp := [26]string{}
	for i, key := range keys {
		mp[key-'a'] = values[i]
	}
	e := Encrypter{mp, map[string]int{}}
	for _, s := range dictionary {
		e.cnt[e.Encrypt(s)]++
	}
	return e
}

func (e *Encrypter) Encrypt(word1 string) string {
	res := make([]byte, 0, len(word1)*2)
	for _, ch := range word1 {
		s := e.mp[ch-'a']
		if s == "" { return "" }
		res = append(res, s...)
	}
	return string(res)
}

func (e *Encrypter) Decrypt(word2 string) int { return e.cnt[word2] }
