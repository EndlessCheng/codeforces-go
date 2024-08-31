package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.codechef.com/problems/VPATH
// https://www.codechef.com/status/vpath?status=Correct
// https://discuss.codechef.com/t/vpath-editorial/
func Test_vpath(t *testing.T) {
	testCases := [][2]string{
		{
			`2
4
1 2
3 1
2 4
4
1 2
2 3
2 4`,
			`15
13`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, vPath)
}
