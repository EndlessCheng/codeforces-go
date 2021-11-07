package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1603/A
// https://codeforces.com/problemset/status/1603/problem/A
func TestCF1603A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
1 2 3
1
2
2
7 7
10
384836991 191890310 576823355 782177068 404011431 818008580 954291757 160449218 155374934 840594328
8
6 69 696 69696 696969 6969696 69696969 696969696
outputCopy
YES
NO
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1603A)
}
