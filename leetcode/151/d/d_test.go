package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	// copy to the Custom Testcase
	const exampleIns = `
["DinnerPlates","push","push","push","push","push","popAtStack","push","push","popAtStack","popAtStack","pop","pop","pop","pop","pop"]
[[2],[1],[2],[3],[4],[5],[0],[20],[21],[0],[2],[],[],[],[],[]]
`
	exampleOuts := `
[null,null,null,null,null,null,2,null,null,20,21,5,4,3,1,-1]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[null,null,null,null,null,null,3,null,null,20,21,5,4,3,1,-1]
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
