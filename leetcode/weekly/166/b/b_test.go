// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`[3,3,3,3,3,1,3]`}, {`[2,1,3,3,3,2]`}}
	exampleOuts := [][]string{{`[[5],[0,1,2],[3,4,6]]`}, {`[[1],[0,5],[2,3,4]]`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, groupThePeople, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}
