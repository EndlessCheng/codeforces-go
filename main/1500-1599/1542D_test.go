package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1542/D
// https://codeforces.com/problemset/status/1542/problem/D
func TestCF1542D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
-
+ 1
+ 2
-
outputCopy
16
inputCopy
15
+ 2432543
-
+ 4567886
+ 65638788
-
+ 578943
-
-
+ 62356680
-
+ 711111
-
+ 998244352
-
-
outputCopy
750759115`
	testutil.AssertEqualCase(t, rawText, 0, CF1542D)
}
