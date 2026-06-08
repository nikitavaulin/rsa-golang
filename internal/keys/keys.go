package keys

import (
	"fmt"
	"math/big"
)

type PublicKey struct {
	E *big.Int
	N *big.Int
}

type SecretKey struct {
	N *big.Int
	D *big.Int
}

func (k *PublicKey) String() string {
	return fmt.Sprintf(
		"e=%s\nn=%s",
		k.E.String(),
		k.N.String(),
	)
}

func (k *SecretKey) String() string {
	return fmt.Sprintf(
		"d=%s\nn=%s",
		k.D.String(),
		k.N.String(),
	)
}
