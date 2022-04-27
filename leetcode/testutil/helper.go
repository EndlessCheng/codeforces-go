package testutil

import (
	"os"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

// 移除每行的左右空格
func trimSpace(s string) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "")
}

// 移除多余空行和左右空格
func trimSpaceAndEmptyLine(s string) (res []string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			res = append(res, line)
		}
	}
	return
}

// IsDebugging will return true if the process was launched from Delve or the
// gopls language server debugger.
//
// It does not detect situations where a debugger attached after process start.
// Adapted from https://stackoverflow.com/a/70969754/4419904
func IsDebugging() bool {
	pid := int32(os.Getppid())

	// We loop in case there were intermediary processes like the gopls language server.
	for pid != 0 {
		p, err := process.NewProcess(pid)
		if err != nil {
			return false
		}
		name, err := p.Name()
		if err != nil {
			return false
		}
		if strings.HasPrefix(name, "dlv") {
			return true
		}
		pid, _ = p.Ppid()
	}
	return false
}
