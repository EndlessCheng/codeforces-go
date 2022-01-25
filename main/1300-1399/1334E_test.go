package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1334/problem/E
// https://codeforces.com/problemset/status/1334/problem/E
func TestCF1334E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
12
3
4 4
12 1
3 4
outputCopy
1
3
1
inputCopy
1
1
1 1
outputCopy
1
inputCopy
288807105787200
4
46 482955026400
12556830686400 897
414 12556830686400
4443186242880 325
outputCopy
547558588
277147129
457421435
702277623`
	testutil.AssertEqualCase(t, rawText, 0, CF1334E)
}
