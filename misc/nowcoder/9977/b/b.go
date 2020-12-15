package main

import (
	"fmt"
	"math/big"
)

// github.com/EndlessCheng/codeforces-go
func judge(n string) bool {
	v := new(big.Int)
	fmt.Sscan(n, v)
	return v.And(v, new(big.Int).Add(v, big.NewInt(1))).Cmp(big.NewInt(0)) == 0
}
