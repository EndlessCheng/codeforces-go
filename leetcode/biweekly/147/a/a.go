package main

import "strings"

// https://space.bilibili.com/206214
func hasMatch(s, p string) bool {
	star := strings.IndexByte(p, '*')
	i := strings.Index(s, p[:star])
	return i >= 0 && strings.Contains(s[i+star:], p[star+1:])
}
