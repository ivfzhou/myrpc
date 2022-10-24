package codec

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const JsonpbSerializerType SerializerType = "jsonpb"

type jsonpbSerializer struct {
	protojson.MarshalOptions
	protojson.UnmarshalOptions
}

func init() {
	RegisterSerializer(JsonpbSerializerType, &jsonpbSerializer{
		protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true, UseProtoNames: true},
		protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true},
	})
}

func (s *jsonpbSerializer) Serialize(v interface{}) ([]byte, error) {
	return s.Marshal(v.(proto.Message))
}

func (s *jsonpbSerializer) Deserialize(data []byte, v interface{}) error {
	return s.Unmarshal(data, v.(proto.Message))
}
