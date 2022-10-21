package codec

type Serializer interface {
	Serialize()
	Deserialize()
}

type Compressor interface {
	Compress()
	Decompress()
}
