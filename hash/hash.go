package hash

type Hash interface {
	// GetAlgorithmName 获取算法名称
	GetAlgorithmName() string

	// WriteContent 写入 字节数据
	WriteContent(buf []byte)

	// CalculateSum 计算 hash 值
	CalculateSum() string
}
