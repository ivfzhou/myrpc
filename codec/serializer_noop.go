package codec

const NoopSerializerType SerializerType = "noop"

type noopSerializer struct{}

func init() {
	RegisterSerializer(NoopSerializerType, &noopSerializer{})
}

func (s *noopSerializer) Serialize(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}
func (s *noopSerializer) Deserialize(data []byte, v interface{}) error {
	copy(v.([]byte), data)
	return nil
}
