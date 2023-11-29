package hash

import (
	"fmt"
	"hash"
	"hash/crc32"
	"strconv"
)

type Crc32 struct {
	table    hash.Hash32
	checkSum uint32
}

func NewCrc32() *Crc32 {
	return &Crc32{table: crc32.NewIEEE()}
}

func (crc Crc32) GetAlgorithmName() string {
	return "crc32"
}

func (crc Crc32) WriteContent(buf []byte) {
	_, err := crc.table.Write(buf)
	if err != nil {
		fmt.Println("crc32 写入字节数据失败")
		return
	}
}

func (crc Crc32) CalculateSum() string {
	sum := int(crc.table.Sum32())
	return fmt.Sprintf("10进制: %s \t | \t 16进制: %s", strconv.Itoa(sum), strconv.FormatInt(int64(sum), 16))

}
