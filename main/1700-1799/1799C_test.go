package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1799/C
// https://codeforces.com/problemset/status/1799/problem/C
func TestCF1799C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
12
a
aab
abb
abc
aabb
aabbb
aaabb
abbb
abbbb
abbcc
eaga
ffcaba
outputCopy
a
aba
bab
bca
abba
abbba
ababa
bbab
bbabb
bbcca
agea
acffba`
	testutil.AssertEqualCase(t, rawText, 0, CF1799C)
}
