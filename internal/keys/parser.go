package keys

import (
	"fmt"
	"math/big"
	"strings"
)

func parseKey(data string, valueKey string) (value *big.Int, modulus *big.Int, err error) {
	lines := strings.Split(data, "\n")
	var valueStr, modulusStr string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, valueKey+"=") {
			valueStr = strings.TrimPrefix(line, valueKey+"=")
		} else if strings.HasPrefix(line, "n=") {
			modulusStr = strings.TrimPrefix(line, "n=")
		}
	}

	if valueStr == "" || modulusStr == "" {
		return nil, nil, fmt.Errorf("invalid key format: missing %s= or n=", valueKey)
	}

	value = new(big.Int)
	modulus = new(big.Int)

	_, ok1 := value.SetString(valueStr, 10)
	_, ok2 := modulus.SetString(modulusStr, 10)

	if !ok1 || !ok2 {
		return nil, nil, fmt.Errorf("failed to parse key values")
	}

	return value, modulus, nil
}

func ParsePublicKey(data string) (PublicKey, error) {
	e, n, err := parseKey(data, "e")
	if err != nil {
		return PublicKey{}, err
	}
	return PublicKey{E: e, N: n}, nil
}

func ParseSecretKey(data string) (SecretKey, error) {
	d, n, err := parseKey(data, "d")
	if err != nil {
		return SecretKey{}, err
	}
	return SecretKey{D: d, N: n}, nil
}
