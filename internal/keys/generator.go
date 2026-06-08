package keys

import (
	"fmt"

	"github.com/nikitavaulin/rsa-golang/internal/rsamath"
)

func GenerateKeys() (PublicKey, SecretKey, error) {
	p, q, err := getSecretPrimeNumber()
	if err != nil {
		return PublicKey{}, SecretKey{}, fmt.Errorf("failed to get secret prime numbers: %w", err)
	}

	n := rsamath.Mul(p, q)
	phi := rsamath.Euler(p, q)

	e, err := getEncryptExponent(phi)
	if err != nil {
		return PublicKey{}, SecretKey{}, fmt.Errorf("failed to get encrypt exponent: %w", err)
	}

	d := rsamath.ModularInverse(e, phi)

	return PublicKey{N: n, E: e}, SecretKey{N: n, D: d}, nil
}
