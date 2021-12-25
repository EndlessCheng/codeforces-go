package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1466/E
// https://codeforces.com/problemset/status/1466/problem/E
func TestCF1466E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
2
1 7
3
1 2 4
4
5 5 5 5
5
6 2 2 1 0
1
0
1
1
6
1 12 123 1234 12345 123456
5
536870912 536870911 1152921504606846975 1152921504606846974 1152921504606846973
outputCopy
128
91
1600
505
0
1
502811676
264880351`
	testutil.AssertEqualCase(t, rawText, 1, CF1466E)
}
