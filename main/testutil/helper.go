package testutil

import (
	"github.com/shirou/gopsutil/v3/process"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func removeExtraSpace(s string) string {
	s = strings.TrimSpace(s)
	sp := strings.Split(s, "\n")
	for i := range sp {
		sp[i] = strings.TrimSpace(sp[i])
	}
	return strings.Join(sp, "\n")
}

const minLenToFile = 750

func handleLongOutput(s string) {
	name := time.Now().Format("150405.000")
	if err := ioutil.WriteFile(name+".txt", []byte(s), 0644); err != nil {
		panic(err)
	}
}

func handleOutput(s string) {
	if len(s) >= minLenToFile {
		handleLongOutput(s)
	}
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

func TransEdges(edges [][2]int) [][]int {
	es := make([][]int, len(edges))
	for i, e := range edges {
		es[i] = []int{e[0], e[1]}
	}
	return es
}
