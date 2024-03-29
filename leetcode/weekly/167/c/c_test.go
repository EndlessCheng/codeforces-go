// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	exampleIns := [][]string{{`[[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]]`, `4`}, {`[[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]]`, `1`}, {`[[1,1,1,1],[1,0,0,0],[1,0,0,0],[1,0,0,0]]`, `6`}, {`[[18,70],[61,1],[25,85],[14,40],[11,96],[97,96],[63,45]]`, `40184`}}
	exampleOuts := [][]string{{`2`}, {`0`}, {`3`}, {`2`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, maxSideLength, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}
