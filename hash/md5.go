package hash

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

type Md5 struct {
	md5 hash.Hash
}

func NewMd5() *Md5 {
	return &Md5{md5: md5.New()}
}

func (md Md5) GetAlgorithmName() string {
	return "md5"
}

func (md Md5) WriteContent(buf []byte) {
	md.md5.Write(buf)
}

func (md Md5) CalculateSum() string {
	sum := md.md5.Sum(nil)
	return hex.EncodeToString(sum)
}
