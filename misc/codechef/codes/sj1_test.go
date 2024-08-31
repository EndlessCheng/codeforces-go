package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.codechef.com/problems/SJ1
// https://www.codechef.com/status/SJ1?status=Correct
// https://discuss.codechef.com/t/SJ1-editorial/
func Test_SJ1(t *testing.T) {
	testCases := [][2]string{
		{
			`1
5
1 2
2 5
1 3
3 4
2 3 4 6 7
1 2 3 2 10`,
			`0 9`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, sj1)
}
