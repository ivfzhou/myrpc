package codec

import "google.golang.org/protobuf/proto"

const ProtobufSerializerType SerializerType = "protobuf"

type protobufSerializer struct{}

func init() {
	RegisterSerializer(ProtobufSerializerType, &protobufSerializer{})
}

func (*protobufSerializer) Serialize(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (*protobufSerializer) Deserialize(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
