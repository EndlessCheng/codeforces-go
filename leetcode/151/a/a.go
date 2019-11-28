package main

import (
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) (ans []string) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	n := len(transactions)
	bad := make([]bool, n)
	for i, ti := range transactions {
		splitsI := strings.Split(ti, ",")
		amount, _ := strconv.Atoi(splitsI[2])
		if amount > 1000 {
			bad[i] = true
		}
		timeI, _ := strconv.Atoi(splitsI[1])
		for j, tj := range transactions {
			splitsJ := strings.Split(tj, ",")
			if splitsJ[0] == splitsI[0] && splitsJ[3] != splitsI[3] {
				timeJ, _ := strconv.Atoi(splitsJ[1])
				if abs(timeI-timeJ) <= 60 {
					bad[i] = true
					bad[j] = true
				}
			}
		}
	}
	for i, b := range bad {
		if b {
			ans = append(ans, transactions[i])
		}
	}
	return
}
