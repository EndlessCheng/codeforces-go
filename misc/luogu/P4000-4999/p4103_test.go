package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_p4103(t *testing.T) {
	cases := [][2]string{
		{
			`10 
2 1 
3 2 
4 1 
5 2 
6 4 
7 5 
8 6 
9 7 
10 9 
5 
2 
5 4 
2
10 4 
2 
5 2 
2
6 1 
2 
6 1`,
			`3 3 3 
6 6 6 
1 1 1 
2 2 2 
2 2 2`,
		},
	}
	tarCase := 0
	testutil.AssertEqualStringCase(t, cases, tarCase, p4103)
}
