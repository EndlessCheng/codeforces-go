package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCompare(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() (a []int) {
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 9)
		a = rg.IntSlice(n, 1, 9)
		return
	}

	runAC := func(a []int) (ans int) {

		return
	}

	// test runAC before run CompareInf
	examples := [][]string{

	}
	if err := testutil.RunLeetCodeFuncWithExamples(t, runAC, examples, 0); err != nil {
		t.Fatal(err)
	}
	return
	testutil.CompareInf(t, inputGenerator, runAC, nil /* fill the func to compare */)
}
