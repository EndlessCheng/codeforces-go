package main

// github.com/EndlessCheng/codeforces-go
func isSumEqual(firstWord, secondWord, targetWord string) bool {
	return str2int(firstWord)+str2int(secondWord) == str2int(targetWord)
}

func str2int(s string) (x int) {
	for _, b := range s {
		x = x*10 + int(b-'a')
	}
	return
}
