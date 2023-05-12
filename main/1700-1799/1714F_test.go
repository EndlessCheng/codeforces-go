package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1714/F
// https://codeforces.com/problemset/status/1714/problem/F
func TestCF1714F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
5 1 2 1
5 2 2 2
5 2 2 3
5 2 2 4
5 3 2 3
4 2 1 1
4 3 1 1
4 1 2 3
7 1 4 1
outputCopy
YES
1 2
4 1
3 1
2 5
YES
4 3
2 5
1 5
5 3
NO
YES
2 4
4 1
2 5
5 3
YES
5 4
4 1
2 5
3 5
YES
2 3
3 4
1 3
NO
YES
4 3
1 2
2 4
NO
inputCopy
8280
3 1 1 1
3 1 1 2
3 1 2 1
3 1 2 2
3 2 1 1
3 2 1 2
3 2 2 1
3 2 2 2
4 1 1 1
4 1 1 2
4 1 1 3
4 1 2 1
4 1 2 2
4 1 2 3
4 1 3 1
4 1 3 2
4 1 3 3
4 2 1 1
4 2 1 2
4 2 1 3
4 2 2 1
4 2 2 2
4 2 2 3
4 2 3 1
4 2 3 2
4 2 3 3
4 3 1 1
4 3 1 2
4 3 1 3
4 3 2 1
4 3 2 2
4 3 2 3
4 3 3 1
4 3 3 2
4 3 3 3
5 1 1 1
5 1 1 2
5 1 1 3
5 1 1 4
5 1 2 1
5 1 2 2
5 1 2 3
5 1 2 4
5 1 3 1
5 1 3 2
5 1 3 3
5 1 3 4
5 1 4 1
5 1 4 2
5 1 4 3
5 1 4 4
5 2 1 1
5 2 1 2
5 2 1 3
5 2 1 4
5 2 2 1
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1714F)
}
