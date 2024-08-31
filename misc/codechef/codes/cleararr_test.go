package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.codechef.com/problems/CLEARARR
// https://www.codechef.com/status/CLEARARR?status=Correct
// https://discuss.codechef.com/t/CLEARARR-editorial/
func Test_clearArr(t *testing.T) {
	testCases := [][2]string{
		{
			`3
5 2 7
9 10 11 12 13
5 0 7
9 9 9 9 9
5 2 7
9 1 2 3 10`,
			`23
45
13`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, clearArr)
}
