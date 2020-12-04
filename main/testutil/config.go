package testutil

import "time"

var (
	// true: test only one case in AssertEqualRunResultsInf / CheckRunResultsInf
	Once bool

	// when DebugTLE > 0, a running case would cause a fatal error when timeout
	DebugTLE = 2 * time.Second
)
