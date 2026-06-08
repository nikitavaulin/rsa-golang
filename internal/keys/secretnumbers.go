package keys

import (
	"fmt"
	"math/big"
	"os"
)

func getSecretPrimeNumber() (p, q *big.Int, err error) {
	pStr, err := getStringEnvVar("p")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get 'p' from .env: %w", err)
	}

	qStr, err := getStringEnvVar("q")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get 'q' from .env: %w", err)
	}

	p, err = parseBigInt(pStr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse 'p' as big integer: %w", err)
	}

	q, err = parseBigInt(qStr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse 'q' as big integer: %w", err)
	}

	return p, q, nil
}

func getEncryptExponent(phi *big.Int) (*big.Int, error) {
	eStr, err := getStringEnvVar("e")
	if err != nil {
		return nil, fmt.Errorf("failed to get 'e' from .env: %w", err)
	}

	e, err := parseBigInt(eStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse 'e' as big integer: %w", err)
	}

	if phi.Uint64() == e.Uint64() {
		return nil, fmt.Errorf("E cannot be equal Phi(N)")
	}

	return e, nil
}

func getStringEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("failed to read '%s' from .env", key)
	}
	return value, nil
}

func parseBigInt(s string) (*big.Int, error) {
	n := new(big.Int)
	_, ok := n.SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("invalid big integer format: %s", s)
	}

	return n, nil
}

// func getNumberEnvVar(key string) (uint64, error) {
// 	value := os.Getenv(key)
// 	if value == "" {
// 		return 0, fmt.Errorf("failed to read '%s' from .env", key)
// 	}

// 	number, err := strconv.ParseUint(value, 10, 64)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to convert env var value to int: %w", err)
// 	}

// 	return number, nil
// }
