package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1509/B
// https://codeforces.com/problemset/status/1509/problem/B
func TestCF1509B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
TMT
3
MTT
6
TMTMTT
6
TMTTTT
6
TTMMTT
outputCopy
YES
NO
YES
NO
YES
inputCopy
3
3
TMT
6
MMTTTT
12
TTTMTTMMMTTT
outputCopy
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1509B)
}
