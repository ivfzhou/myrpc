package codec

import jsoniter "github.com/json-iterator/go"

const JsoniterSerializerType SerializerType = "jsoniter"

type jsoniterSerializer struct {
	jsoniter.API
}

func init() {
	RegisterSerializer(JsoniterSerializerType, &jsoniterSerializer{jsoniter.ConfigCompatibleWithStandardLibrary})
}

func (s *jsoniterSerializer) Serialize(v interface{}) ([]byte, error) {
	return s.Marshal(v)
}

func (s *jsoniterSerializer) Deserialize(data []byte, v interface{}) error {
	return s.Unmarshal(data, v)
}
