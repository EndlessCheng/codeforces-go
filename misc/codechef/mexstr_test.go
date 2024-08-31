package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.codechef.com/problems/MEXSTR
// https://www.codechef.com/status/MEXSTR?status=Correct
// https://discuss.codechef.com/t/mexstr-editorial/
func Test_mexStr(t *testing.T) {
	testCases := [][2]string{
		{
			`2
1001011
1111`,
			`1100
0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, mexStr)
}
