package codec

import "encoding/xml"

const XMLSerializerType SerializerType = "xml"

type xmlSerializer struct{}

func init() {
	RegisterSerializer(XMLSerializerType, &xmlSerializer{})
}

func (*xmlSerializer) Serialize(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (*xmlSerializer) Deserialize(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}
