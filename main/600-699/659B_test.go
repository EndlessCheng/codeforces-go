package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/659/B
// https://codeforces.com/problemset/status/659/problem/B
func TestCF659B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
Ivanov 1 763
Andreev 2 800
Petrov 1 595
Sidorov 1 790
Semenov 2 503
outputCopy
Sidorov Ivanov
Andreev Semenov
inputCopy
5 2
Ivanov 1 800
Andreev 2 763
Petrov 1 800
Sidorov 1 800
Semenov 2 503
outputCopy
?
Andreev Semenov`
	testutil.AssertEqualCase(t, rawText, 0, CF659B)
}
