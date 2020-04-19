package copypasta

import "math/big"

type bigInt struct{ *big.Int }

func (a bigInt) add(b bigInt) bigInt { a.Add(a.Int, b.Int); return a }
func (a bigInt) sub(b bigInt) bigInt { a.Sub(a.Int, b.Int); return a }
func (a bigInt) mul(b bigInt) bigInt { a.Mul(a.Int, b.Int); return a }
func (a bigInt) div(b bigInt) bigInt { a.Quo(a.Int, b.Int); return a }
func (a bigInt) mod(b bigInt) bigInt { a.Rem(a.Int, b.Int); return a }
func (a bigInt) neg() bigInt         { a.Neg(a.Int); return a }
func (a bigInt) cmp(b bigInt) int    { return a.Cmp(b.Int) }

type bigRat struct{ *big.Rat }

func (a bigRat) add(b bigRat) bigRat { a.Add(a.Rat, b.Rat); return a }
func (a bigRat) sub(b bigRat) bigRat { a.Sub(a.Rat, b.Rat); return a }
func (a bigRat) mul(b bigRat) bigRat { a.Mul(a.Rat, b.Rat); return a }
func (a bigRat) div(b bigRat) bigRat { a.Quo(a.Rat, b.Rat); return a }
func (a bigRat) neg() bigRat         { a.Neg(a.Rat); return a }
func (a bigRat) cmp(b bigRat) int    { return a.Cmp(b.Rat) }
