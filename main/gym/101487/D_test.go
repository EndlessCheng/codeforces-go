package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/gym/101487/problem/D
// https://codeforces.com/gym/101487/status/D
func Test_runD(t *testing.T) {
	testCases := [][2]string{
		{
			`2
9
5 3 4 9 2 8 6 7 1
7
1 2 3 10 4 5 6`,
			`4
6`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, runD)
}
