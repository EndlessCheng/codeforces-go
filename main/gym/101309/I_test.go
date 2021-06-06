package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/gym/101309
func Test_runI(t *testing.T) {
	testCases := [][2]string{
		{
			`4 6
1 2 1
1 3 2
3 4 3
2 3 1
2 4 4
3 1 1`,
			`2
1 3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, runI)
}