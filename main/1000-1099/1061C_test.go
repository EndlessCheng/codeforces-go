package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1061/C
// https://codeforces.com/problemset/status/1061/problem/C
func TestCF1061C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
3
inputCopy
5
2 2 1 22 14
outputCopy
13
inputCopy
15
513046 683844 914823 764255 815301 790234 184972 93547 388028 211665 554415 713159 183950 200951 842336
outputCopy
161`
	testutil.AssertEqualCase(t, rawText, 0, CF1061C)
}
