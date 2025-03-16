package main

import (
	"strconv"
	"strings"
	"unicode"
)

// https://space.bilibili.com/206214
type Spreadsheet map[string]int

func Constructor(int) Spreadsheet {
	return Spreadsheet{}
}

func (s Spreadsheet) SetCell(cell string, value int) {
	s[cell] = value
}

func (s Spreadsheet) ResetCell(cell string) {
	delete(s, cell)
}

func (s Spreadsheet) GetValue(formula string) (ans int) {
	for _, cell := range strings.Split(formula[1:], "+") {
		if unicode.IsUpper(rune(cell[0])) {
			ans += s[cell]
		} else {
			x, _ := strconv.Atoi(cell)
			ans += x
		}
	}
	return
}
