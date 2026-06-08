package rsacrypto

import (
	"math/big"
	"strings"

	"github.com/nikitavaulin/rsa-golang/internal/converter"
	"github.com/nikitavaulin/rsa-golang/internal/keys"
	"github.com/nikitavaulin/rsa-golang/internal/rsamath"
)

func Encrypt(text []byte, key keys.PublicKey) []byte {
	blockSize := calcBlockSize(key.N)

	binMsg := getBinaryMsgPresentation(text, blockSize)

	var binCipher strings.Builder
	cipherBlockSize := blockSize + 1

	for i := 0; i < len(binMsg); i += blockSize {
		block := binMsg[i : i+blockSize]

		m := converter.BinToDecBigInt(block)
		c := encryptBlock(m, key)

		cBin := converter.BigIntToBinFullSize(c, cipherBlockSize)
		binCipher.WriteString(cBin)
	}

	cipher := []byte(binCipher.String())
	return cipher
}

func encryptBlock(m *big.Int, key keys.PublicKey) *big.Int {
	return rsamath.PowerMod(m, key.E, key.N)
}
