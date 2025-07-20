package mirrorbase

import (
	"github.com/v4fly/v4ray-core/v0/common/crypto"
)

func reverseBytesGeneratorByteOrder(generator crypto.BytesGenerator) crypto.BytesGenerator {
	var reverseResult [8]byte
	return func() []byte {
		result := generator()
		if len(result) != 8 {
			panic("reverseBytesGeneratorByteOrder requires a generator that returns exactly 8 bytes")
		}
		for i := 0; i < 8; i++ {
			reverseResult[i] = result[7-i]
		}
		return reverseResult[:]
	}
}
