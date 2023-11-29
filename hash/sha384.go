package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type Sha384 struct {
	sha384 hash.Hash
}

func NewSha384() *Sha384 {
	return &Sha384{
		sha384: sha512.New384(),
	}
}

func (sha Sha384) GetAlgorithmName() string {
	return "sha384"
}

func (sha Sha384) WriteContent(buf []byte) {
	sha.sha384.Write(buf)
}

func (sha Sha384) CalculateSum() string {
	sum := sha.sha384.Sum(nil)
	return hex.EncodeToString(sum)
}
