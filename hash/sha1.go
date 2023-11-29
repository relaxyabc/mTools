package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

type Sha1 struct {
	sha1 hash.Hash
}

func NewSha1() *Sha1 {
	return &Sha1{sha1: sha1.New()}
}

func (sha Sha1) GetAlgorithmName() string {
	return "sha1"
}

func (sha Sha1) WriteContent(buf []byte) {
	sha.sha1.Write(buf)
}

func (sha Sha1) CalculateSum() string {
	sum := sha.sha1.Sum(nil)
	return hex.EncodeToString(sum)
}
