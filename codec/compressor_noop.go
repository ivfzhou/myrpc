package codec

const NoopCompressorType CompressorType = "noop"

type noopCompressor struct{}

func init() {
	RegisterCompressor(NoopCompressorType, &noopCompressor{})
}

func (*noopCompressor) Compress(data []byte) ([]byte, error) {
	return data, nil
}

func (*noopCompressor) Decompress(data []byte) ([]byte, error) {
	return data, nil
}
