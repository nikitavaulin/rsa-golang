package rsacrypto

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/nikitavaulin/rsa-golang/internal/converter"
)

func calcBlockSize(n *big.Int) int {
	return n.BitLen() - 1
}

func getBinaryMsgPresentation(data []byte, blockSize int) string {
	binSeq := ""
	for _, b := range data {
		bin := fmt.Sprintf("%b", b)
		bin = converter.AddInsignificantZeros(bin, 8)

		binSeq += bin
	}

	blocksCount := (len(binSeq) + blockSize - 1) / blockSize
	expectedBinSeqLength := blocksCount * blockSize
	binSeq = converter.AddInsignificantZeros(binSeq, expectedBinSeqLength)

	return binSeq
}

func getPlainTextBytePresentation(binText string) []byte {
	binTextLen := len(binText)
	if binTextLen == 0 {
		return nil
	}

	blocksCount := binTextLen / 8
	result := make([]byte, blocksCount)

	targetMsgLen := 8 * blocksCount
	insignZerosCount := binTextLen - targetMsgLen

	byteIdx := 0
	for i := insignZerosCount; i < binTextLen; i += 8 {
		block := binText[i : i+8]
		b, _ := strconv.ParseUint(block, 2, 8)
		// fmt.Println(block, b)
		result[byteIdx] = byte(b)
		byteIdx++
	}

	return result
}

// func getCipherBytePresentation(binCipher string) []byte {
// 	binCipherLen := len(binCipher)
// 	if binCipherLen == 0 {
// 		return nil
// 	}

// 	blocksCount := (binCipherLen + 7) / 8
// 	result := make([]byte, blocksCount)

// 	targetLen := 8 * blocksCount
// 	binCipher = converter.AddInsignificantZeros(binCipher, targetLen)

// 	byteIdx := 0
// 	for i := 0; i < binCipherLen; i += 8 {
// 		block := binCipher[i : i+8]
// 		b, _ := strconv.ParseUint(block, 2, 8)

// 		result[byteIdx] = byte(b)
// 		byteIdx++
// 	}

// 	return result
// }
