package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	// copy to the Custom Testcase
	const exampleIns = `
[1,7,0,7,-8,null,null]
`
	exampleOuts := `
2
`
	// copy Your answer in the Run Code Result
	yourAnswers := `
2
`
	assert.Equal(t, strings.TrimSpace(exampleOuts), strings.TrimSpace(yourAnswers))
}
