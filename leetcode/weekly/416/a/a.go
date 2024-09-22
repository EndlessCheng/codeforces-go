package main

// https://space.bilibili.com/206214
func reportSpam(message, bannedWords []string) bool {
	banned := map[string]bool{}
	for _, s := range bannedWords {
		banned[s] = true
	}
	seen := false
	for _, s := range message {
		if banned[s] {
			if seen {
				return true
			}
			seen = true
		}
	}
	return false
}
