// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`[1,2,3,3,4,4,5,6]`, `4`}, {`[3,2,1,2,3,4,3,4,5,9,10,11]`, `3`}, {`[3,3,2,2,1,1]`, `3`}, {`[1,2,3,4]`, `3`}}
	exampleOuts := [][]string{{`true`}, {`true`}, {`true`}, {`false`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, isPossibleDivide, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}
