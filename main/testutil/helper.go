package testutil

import (
	"io/ioutil"
	"time"
)

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
