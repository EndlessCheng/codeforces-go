package copypasta

import "math/big"

// CF1244C

// a+=b: a.add(b)
// a+b:  add(a,b)
type Int struct{ *big.Int }

func newInt(x int64) Int    { return Int{big.NewInt(x)} }
func (a Int) add(b Int) Int { a.Add(a.Int, b.Int); return a }
func (a Int) sub(b Int) Int { a.Sub(a.Int, b.Int); return a }
func (a Int) mul(b Int) Int { a.Mul(a.Int, b.Int); return a }
func (a Int) div(b Int) Int { a.Quo(a.Int, b.Int); return a }
func (a Int) mod(b Int) Int { a.Rem(a.Int, b.Int); return a }
func (a Int) neg() Int      { a.Neg(a.Int); return a }
func (a Int) cmp(b Int) int { return a.Cmp(b.Int) }

func add(a, b Int) Int { return Int{(&big.Int{}).Add(a.Int, b.Int)} }
func sub(a, b Int) Int { return Int{(&big.Int{}).Sub(a.Int, b.Int)} }
func mul(a, b Int) Int { return Int{(&big.Int{}).Mul(a.Int, b.Int)} }
func div(a, b Int) Int { return Int{(&big.Int{}).Quo(a.Int, b.Int)} }
func mod(a, b Int) Int { return Int{(&big.Int{}).Rem(a.Int, b.Int)} }
func neg(a Int) Int    { return Int{(&big.Int{}).Neg(a.Int)} }

//

type rat struct{ *big.Rat }

func newRat(a, b int64) rat { return rat{big.NewRat(a, b)} }
func (a rat) add(b rat) rat { a.Add(a.Rat, b.Rat); return a }
func (a rat) sub(b rat) rat { a.Sub(a.Rat, b.Rat); return a }
func (a rat) mul(b rat) rat { a.Mul(a.Rat, b.Rat); return a }
func (a rat) div(b rat) rat { a.Quo(a.Rat, b.Rat); return a }
func (a rat) neg() rat      { a.Neg(a.Rat); return a }
func (a rat) cmp(b rat) int { return a.Cmp(b.Rat) }
