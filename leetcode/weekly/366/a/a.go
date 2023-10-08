package main

// https://space.bilibili.com/206214
func differenceOfSums(n, m int) int {
	return n*(n+1)/2 - n/m*(n/m+1)*m
}
