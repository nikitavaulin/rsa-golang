package rsacrypto

import (
	"math/big"
	"strings"

	"github.com/nikitavaulin/rsa-golang/internal/converter"
	"github.com/nikitavaulin/rsa-golang/internal/keys"
	"github.com/nikitavaulin/rsa-golang/internal/rsamath"
)

func Decrypt(cipher []byte, key keys.SecretKey) []byte {
	blockSize := calcBlockSize(key.N)
	cipherBlockSize := blockSize + 1

	cipherLen := len(cipher)

	blocksCount := cipherLen / cipherBlockSize
	realCipherLength := cipherBlockSize * blocksCount
	start := cipherLen - realCipherLength

	var decryptedTextBin strings.Builder

	for i := start; i < cipherLen; i += cipherBlockSize {
		block := cipher[i : i+cipherBlockSize]

		c := converter.BinToDecBigInt(string(block))
		m := decryptBlock(c, key)

		msgBin := converter.BigIntToBinFullSize(m, blockSize)

		decryptedTextBin.WriteString(msgBin)
	}

	decryptedText := getPlainTextBytePresentation(decryptedTextBin.String())
	return decryptedText
}

func decryptBlock(c *big.Int, key keys.SecretKey) *big.Int {
	return rsamath.PowerMod(c, key.D, key.N)
}
