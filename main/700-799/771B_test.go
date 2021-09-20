package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/771/B
// https://codeforces.com/problemset/status/771/problem/B
func TestCF771B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 3
NO NO YES YES YES NO
outputCopy
Adam Bob Bob Cpqepqwer Limak Adam Bob Adam
inputCopy
9 8
YES NO
outputCopy
R Q Ccccccccc Ccocc Ccc So Strong Samples Ccc
inputCopy
3 2
NO NO
outputCopy
Na Na Na`
	testutil.AssertEqualCase(t, rawText, 0, CF771B)
}
