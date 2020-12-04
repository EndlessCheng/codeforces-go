package testutil

import "time"

var (
	// 方便打断点，配合 targetCaseNum 一起使用
	DebugCallIndex int

	DebugTLE = 2 * time.Second
)
