package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`4 5 
1 2 1 
1 3 3 
1 4 1 
2 3 4 
3 4 1 `,
			`4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
