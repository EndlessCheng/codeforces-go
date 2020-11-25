package main

import (
	"fmt"
	"math/big"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Maxsumforknumers(S string, k int) string {
	s := []byte(S)
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	sum, n := 0, len(S)
	for _, b := range s[n-k+1:] {
		sum += int(b & 15)
	}
	v := new(big.Int)
	fmt.Sscan(string(s[:n-k+1]), v)
	return v.Add(v, big.NewInt(int64(sum))).String()
}
