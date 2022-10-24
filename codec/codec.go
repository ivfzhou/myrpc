package codec

import "sync"

var (
	serializerMap sync.Map
	compressorMap sync.Map
)

type Serializer interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) error
}

type Compressor interface {
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}

type (
	SerializerType string
	CompressorType string
)

func GetSerializer(name SerializerType) Serializer {
	value, ok := serializerMap.Load(name)
	if ok && value != nil {
		return value.(Serializer)
	}
	return nil
}

func GetCompressor(name CompressorType) Compressor {
	value, ok := compressorMap.Load(name)
	if ok && value != nil {
		return value.(Compressor)
	}
	return nil
}

func RegisterSerializer(name SerializerType, serializer Serializer) {
	serializerMap.Store(name, serializer)
}

func RegisterCompressor(name CompressorType, compressor Compressor) {
	compressorMap.Store(name, compressor)
}

func DeregisterSerializer(name SerializerType) {
	serializerMap.Delete(name)
}

func DeregisterCompressor(name CompressorType) {
	compressorMap.Delete(name)
}
