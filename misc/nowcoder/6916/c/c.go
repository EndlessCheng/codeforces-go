package main

import (
	//. "nc_tools"
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// github.com/EndlessCheng/codeforces-go
const M int = 1e9 + 7

func pow(x int) int {
	p := 1
	for n := M - 2; n > 0; n >>= 1 {
		if n&1 == 1 {
			p = p * x % M
		}
		x = x * x % M
	}
	return p
}

func NTimesPoint(p *Point, n int) *Point {
	if n == 1 {
		return p
	}
	q := *NTimesPoint(p, n>>1)
	k := (3*q.X*q.X + 1) % M * pow(2*q.Y) % M
	x := (k*k + 2*(M-q.X)) % M
	q.X, q.Y = x, ((q.X-x+M)*k-q.Y+M)%M
	if n&1 > 0 {
		k = (p.Y - q.Y + M) * pow(p.X-q.X+M) % M
		x = (k*k - p.X - q.X + 2*M) % M
		q.X, q.Y = x, ((q.X-x+M)*k-q.Y+M)%M
	}
	return &q
}
