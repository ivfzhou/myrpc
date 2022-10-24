package codec

import "encoding/json"

const JSONSerializerType SerializerType = "json"

type jsonSerializer struct{}

func init() {
	RegisterSerializer(JSONSerializerType, &jsonSerializer{})
}

func (*jsonSerializer) Serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*jsonSerializer) Deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
