package codec

import (
	"bytes"
	"compress/gzip"
	"io"
	"sync"
)

const GzipCompressorType CompressorType = "gzip"

type gzipCompressor struct {
	readers sync.Pool
	writers sync.Pool
}

func init() {
	RegisterCompressor(GzipCompressorType, &gzipCompressor{})
}

func (c *gzipCompressor) Compress(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	writer, ok := c.writers.Get().(*gzip.Writer)
	if ok {
		writer.Reset(buf)
	} else {
		writer = gzip.NewWriter(buf)
	}
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

func (c *gzipCompressor) Decompress(data []byte) ([]byte, error) {
	buf := bytes.NewReader(data)
	reader, ok := c.readers.Get().(*gzip.Reader)
	var err error
	if ok {
		err = reader.Reset(buf)
	} else {
		reader, err = gzip.NewReader(buf)
	}
	if err != nil {
		return nil, err
	}
	defer c.readers.Put(reader)
	return io.ReadAll(reader)
}
