package codec

import (
	"bytes"
	"compress/lzw"
	"io"
	"sync"
)

const LzwCompressorType CompressorType = "lzw"

type lzwCompressor struct {
	readers sync.Pool
	writers sync.Pool
}

func init() {
	RegisterCompressor(LzwCompressorType, &lzwCompressor{})
}

func (c *lzwCompressor) Compress(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	writer, ok := c.writers.Get().(*lzw.Writer)
	if !ok {
		writer = &lzw.Writer{}
	}
	writer.Reset(buf, lzw.MSB, 8)
	defer c.writers.Put(writer)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *lzwCompressor) Decompress(data []byte) ([]byte, error) {
	buf := bytes.NewReader(data)
	reader, ok := c.readers.Get().(*lzw.Reader)
	if !ok {
		reader = &lzw.Reader{}
	}
	reader.Reset(buf, lzw.MSB, 8)
	defer c.readers.Put(reader)
	return io.ReadAll(reader)
}
