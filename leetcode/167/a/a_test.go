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
[1,0,1]
[0]
[1]
[1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
[0,0]
`
	exampleOuts := `
5
0
1
18880
0
`
	// copy Your answer in the Run Code Result
	yourAnswers := `

`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
