package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.codechef.com/problems/CONDEL
// https://www.codechef.com/status/CONDEL?status=Correct
// https://discuss.codechef.com/t/CONDEL-editorial/
func Test_conDel(t *testing.T) {
	testCases := [][2]string{
		{
			`3
5 3
1 0 1 0 1
4 4
1 1 0 1
3 1
1 1 0`,
			`3
6
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, conDel)
}
