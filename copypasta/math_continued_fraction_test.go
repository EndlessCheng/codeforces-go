package copypasta

import (
	"math"
	"math/big"
	"testing"
)

func Test_calcSqrtRatContinuedFraction(t *testing.T) {
	calcGCD := func(a, b int64) int64 {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	calcGCDN := func(nums ...int64) (gcd int64) {
		gcd = nums[0]
		for _, v := range nums[1:] {
			gcd = calcGCD(gcd, v)
		}
		return
	}
	// sqrt(m/n)
	calcSqrtRatContinuedFraction := func(m, n int64) (exp []int64) {
		sqrtRat := math.Sqrt(float64(m) / float64(n))
		base0 := int64(sqrtRat)
		if base0*base0*n == m {
			return []int64{base0}
		}
		a := []int64{1}
		b := []int64{0}
		c := []int64{1}
		const loop = 50
		for i := 0; i < loop; i++ {
			base := int64((float64(a[i])*sqrtRat + float64(b[i])) / float64(c[i]))
			exp = append(exp, base)
			tmp := base*c[i] - b[i]
			newA := n * c[i] * a[i]
			newB := n * c[i] * tmp
			newC := m*a[i]*a[i] - n*tmp*tmp
			if newC == 0 {
				return
			}
			gcd := calcGCDN(newA, newB, newC)
			//Println(i, base, newA, newB, newC, gcd)
			a = append(a, newA/gcd)
			b = append(b, newB/gcd)
			c = append(c, newC/gcd)
		}
		return
	}
	calcSqrtRatContinuedFractionBig := func(m, n int64) (exp []*big.Int) {
		sqrtRat := new(big.Float).Sqrt(big.NewFloat(float64(m) / float64(n)))
		if base0, acc := sqrtRat.Int(nil); acc == big.Exact {
			exp = append(exp, base0)
			return
		}
		bigM, bigN := big.NewInt(m), big.NewInt(n)
		a := []*big.Int{big.NewInt(1)}
		b := []*big.Int{big.NewInt(0)}
		c := []*big.Int{big.NewInt(1)}
		const loop = 20
		for i := 0; i < loop; i++ {
			tmpF := new(big.Float).SetInt(a[i])
			base, _ := tmpF.Mul(tmpF, sqrtRat).
				Add(tmpF, new(big.Float).SetInt(b[i])).
				Quo(tmpF, new(big.Float).SetInt(c[i])).
				Int(nil)
			exp = append(exp, new(big.Int).Set(base))
			tmp := &big.Int{}
			tmp.Mul(base, c[i]).Sub(tmp, b[i])
			newA, newB, newC := &big.Int{}, &big.Int{}, &big.Int{}
			newA.Mul(bigN, c[i]).Mul(newA, a[i])
			newB.Mul(bigN, c[i]).Mul(newB, tmp)
			tmp.Mul(tmp, tmp).Mul(tmp, bigN)
			newC.Mul(bigM, a[i]).Mul(newC, a[i]).Sub(newC, tmp)
			if newC.IsInt64() && newC.Int64() == 0 {
				return
			}
			gcd := big.NewInt(1)
			if !base.IsInt64() || base.Int64() > 0 {
				gcd.GCD(nil, nil, newA, newB).GCD(nil, nil, gcd, newC)
			}
			//Println(i, base, newA, newB, newC, gcd)
			a = append(a, newA.Quo(newA, gcd))
			b = append(b, newB.Quo(newB, gcd))
			c = append(c, newC.Quo(newC, gcd))
		}
		return
	}

	t.Log(calcSqrtRatContinuedFraction(3, 10))
	t.Log(calcSqrtRatContinuedFractionBig(3, 10))
	t.Log(calcSqrtRatContinuedFraction(4, 1))
	t.Log(calcSqrtRatContinuedFractionBig(4, 1))
	t.Log(calcSqrtRatContinuedFraction(1, 4))
	t.Log(calcSqrtRatContinuedFractionBig(1, 4))

	t.Log(calcSqrtRatContinuedFractionBig(1, 1e9))
	t.Log(calcSqrtRatContinuedFractionBig(2, 1e9))
	t.Log(calcSqrtRatContinuedFractionBig(3, 1e9))
	t.Log(calcSqrtRatContinuedFractionBig(1e9-1, 1e9))
	t.Log(calcSqrtRatContinuedFractionBig(1e9-2, 1e9))
	t.Log(calcSqrtRatContinuedFractionBig(1e9-3, 1e9))
	t.Log(calcSqrtRatContinuedFractionBig(1e9-4, 1e9))
	t.Log(calcSqrtRatContinuedFractionBig(1e9-5, 1e9))
}
