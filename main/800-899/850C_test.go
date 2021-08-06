package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/850/C
// https://codeforces.com/problemset/status/850/problem/C
func TestCF850C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 1 1
outputCopy
Arpa
inputCopy
4
1 1 17 17
outputCopy
Mojtaba
inputCopy
4
1 1 17 289
outputCopy
Arpa
inputCopy
5
1 2 3 4 5
outputCopy
Arpa
inputCopy
6
9 5 1 1 8 1
outputCopy
Arpa`
	testutil.AssertEqualCase(t, rawText, -1, CF850C)
}
