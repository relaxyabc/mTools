package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type Sha256 struct {
	sha256 hash.Hash
}

func NewSha256() *Sha256 {
	return &Sha256{sha256: sha256.New()}
}

func (sha Sha256) GetAlgorithmName() string {
	return "sha256"
}

func (sha Sha256) WriteContent(buf []byte) {
	sha.sha256.Write(buf)
}

func (sha Sha256) CalculateSum() string {
	sum := sha.sha256.Sum(nil)
	return hex.EncodeToString(sum)
}
