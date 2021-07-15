package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/gym/100723/problem/B
// https://codeforces.com/gym/100723/status/B
func Test_runB(t *testing.T) {
	testCases := [][2]string{
		{
			`2
5 3.5
1 1 1 1
2 3 0 1
3 5 1 1
5 1 1 1
5 4 0 1
3 1.1
-1 0 5 10
0 0 3 9
2 0 1 1`,
			`1 2 4
-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, runB)
}
