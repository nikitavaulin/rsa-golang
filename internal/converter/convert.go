package converter

import (
	"math/big"
	"strings"
)

func BinToDecBigInt(bin string) *big.Int {
	number := new(big.Int)
	number.SetString(bin, 2)
	return number
}

func BigIntToBinFullSize(number *big.Int, targetLength int) string {
	bin := BigIntToBin(number)
	bin = AddInsignificantZeros(bin, targetLength)
	return bin
}

func BigIntToBin(number *big.Int) string {
	return number.Text(2)
}

func ByteToBigInt(b byte) *big.Int {
	return big.NewInt(int64(b))
}

func AddInsignificantZeros(binNumber string, targetLength int) string {
	numberLength := len(binNumber)
	if numberLength < targetLength {
		binNumber = strings.Repeat("0", targetLength-numberLength) + binNumber
	}
	return binNumber
}
