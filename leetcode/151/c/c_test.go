package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	// copy to the Custom Testcase
	const exampleIns = `
[1,2,-3,3,1]
[1,2,3,-3,4]
[1,2,3,-3,-2]
`
	exampleOuts := `
[3,1]
[1,2,4]
[1]
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
[3,1]
[1,2,4]
[1]
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
