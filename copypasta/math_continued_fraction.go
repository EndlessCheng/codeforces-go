package copypasta

import (
	"math"
	"math/big"
)

func continuedFractionCollections() {
	// a/b = [exp[0]; exp[1],...]
	calcContinuedFraction := func(a, b int64) (exp []int64) {
		for b != 1 {
			exp = append(exp, a/b)
			a, b = b, a%b
		}
		exp = append(exp, a)
		return
	}
	calcContinuedFraction = func(a_, b_ int64) (exp []int64) {
		a, b := big.NewInt(a_), big.NewInt(b_)
		for !b.IsInt64() || b.Int64() != 1 {
			r := &big.Int{}
			a.QuoRem(a, b, r)
			exp = append(exp, a.Int64())
			a, b = b, r
		}
		exp = append(exp, a.Int64())
		return
	}

	// sqrt(d) = [exp[0]; exp[1],..., 2*exp[0], exp[1], ..., 2*exp[0], exp[1], ...]
	// https://en.wikipedia.org/wiki/Pell%27s_equation
	// https://www.weiwen.io/post/about-the-pell-equations-2/
	calcSqrtContinuedFraction := func(d int64) (exp []int64) {
		sqrtD := math.Sqrt(float64(d))
		base := int64(sqrtD)
		if base*base == d {
			return []int64{base}
		}
		p := []int64{0}
		q := []int64{1}
		for i := 0; ; i++ {
			a := int64((float64(p[i]) + sqrtD) / float64(q[i]))
			exp = append(exp, a)
			if a == 2*base {
				break
			}
			pi := a*q[i] - p[i]
			p = append(p, pi)
			q = append(q, (d-pi*pi)/q[i])
		}
		return
	}

	calcRatByContinuedFraction := func(exp []int64) (a, b int64) {
		n := len(exp)
		h := make([]int64, n+2)
		h[0], h[1] = 0, 1
		k := make([]int64, n+2)
		k[0], k[1] = 1, 0
		for i, a := range exp {
			h[i+2] = a*h[i+1] + h[i]
			k[i+2] = a*k[i+1] + k[i]
		}
		return h[n+1], k[n+1]
	}

	_ = []interface{}{calcContinuedFraction, calcSqrtContinuedFraction, calcRatByContinuedFraction}
}
