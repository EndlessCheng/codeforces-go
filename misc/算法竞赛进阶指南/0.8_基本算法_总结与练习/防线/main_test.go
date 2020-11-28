package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test(t *testing.T) {
	testCases := [][2]string{
		{
			`3
2
1 10 1
2 10 1
2
1 10 1 
1 10 1 
4
1 10 1 
4 4 1 
1 5 1 
6 10 1`,
			`1 1
There's no weakness.
4 3`,
		},
		{
			`3
4
1 10 1
4 4 1
1 5 1
6 10 1
2
1 10 1
2 10 1
2
1 10 1
1 10 1`,
			`4 3
1 1
There's no weakness.`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
