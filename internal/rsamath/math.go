package rsamath

import (
	"math/big"
)

func Euler(p, q *big.Int) *big.Int {
	pMinus1 := new(big.Int).Sub(p, big.NewInt(1))
	qMinus1 := new(big.Int).Sub(q, big.NewInt(1))

	phi := new(big.Int).Mul(pMinus1, qMinus1)
	return phi
}

func GCD(a, b *big.Int) *big.Int {
	return new(big.Int).GCD(nil, nil, a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func PowerMod(base, exp, mod *big.Int) *big.Int {
	return new(big.Int).Exp(base, exp, mod)
}

func ModularInverse(a, mod *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, mod)
}

// func GetPrimeBigInt() *big.Int {
// 	return big.NewInt(65537)
// }
