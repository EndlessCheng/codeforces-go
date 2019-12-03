package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	// copy to the Custom Testcase
	const sampleIns = `
["FindElements","find","find"]
[[[-1,null,-1]],[1],[2]]
["FindElements","find","find","find"]
[[[-1,-1,-1,-1,-1]],[1],[3],[5]]
["FindElements","find","find","find","find"]
[[[-1,null,-1,-1,null,-1]],[2],[3],[4],[5]]
`
	sampleOuts := `
[null,false,true]
[null,true,true,false]
[null,true,false,false,true]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[null,false,true]
[null,true,true,false]
[null,true,false,false,true]
`
	assert.Equal(t, strings.TrimSpace(sampleOuts), strings.TrimSpace(yourAnswers))
}
