package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/467/B
// https://codeforces.com/problemset/status/467/problem/B
func TestCF467B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3 1
8
5
111
17
outputCopy
0
inputCopy
3 3 3
1
2
3
4
outputCopy
3
inputCopy
6 8 2
46
59
38
5
13
54
26
62
18
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF467B)
}
