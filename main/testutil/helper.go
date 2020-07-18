package testutil

import (
	"io/ioutil"
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
