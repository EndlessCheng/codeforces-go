package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/920/G
// https://codeforces.com/problemset/status/920/problem/G
func TestCF920G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7 22 1
7 22 2
7 22 3
outputCopy
9
13
15
inputCopy
5
42 42 42
43 43 43
44 44 44
45 45 45
46 46 46
outputCopy
187
87
139
128
141`
	testutil.AssertEqualCase(t, rawText, 0, CF920G)
}
