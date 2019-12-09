package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	// copy to the Custom Testcase
	const exampleIns = `
1
5
2
5
`
	exampleOuts := `
[[1,4],[2,3],[3,2],[4,1]]
[[1,5],[5,1]]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[[1,4],[2,3],[3,2],[4,1]]
[[1,5],[5,1]]
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
