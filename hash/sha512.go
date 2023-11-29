package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type Sha512 struct {
	sha512 hash.Hash
}

func NewSha512() *Sha512 {
	return &Sha512{sha512: sha512.New()}
}

func (sha Sha512) GetAlgorithmName() string {
	return "sha512"
}

func (sha Sha512) WriteContent(buf []byte) {
	sha.sha512.Write(buf)
}

func (sha Sha512) CalculateSum() string {
	sum := sha.sha512.Sum(nil)
	return hex.EncodeToString(sum)
}
