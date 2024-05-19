package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://space.bilibili.com/206214
func cf1028B(in io.Reader, out io.Writer) {
	Fprintln(out, strings.Repeat("5", 300))
	Fprintln(out, strings.Repeat("4", 299)+"5")
}

//func main() { cf1028B(os.Stdin, os.Stdout) }
