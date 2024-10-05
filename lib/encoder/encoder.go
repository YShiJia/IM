/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 11:13:28
 */

package encoder

import (
	"encoding/json"
	"errors"
	"google.golang.org/protobuf/proto"
)

var ErrMessageTypeNotSupportEncode = errors.New("[Encoder]: message type not support encode")

type Encoder interface {
	Encode(any) ([]byte, error)
	// Decode 传入any，需要是指针类型
	Decode([]byte, any) error
}

type protobufEncoder struct{}

func NewProtobufEncoder() *protobufEncoder {
	return &protobufEncoder{}
}

func (p *protobufEncoder) Encode(message any) ([]byte, error) {
	msg, ok := message.(proto.Message)
	if !ok {
		return nil, ErrMessageTypeNotSupportEncode
	}
	return proto.Marshal(msg)
}

func (p *protobufEncoder) Decode(data []byte, entity any) error {
	e, ok := entity.(proto.Message)
	if !ok {
		return ErrMessageTypeNotSupportEncode
	}
	return proto.Unmarshal(data, e)
}

type jsonEncoder struct{}

func NewJsonEncoder() *jsonEncoder {
	return &jsonEncoder{}
}

func (j *jsonEncoder) Encode(message any) ([]byte, error) {
	return json.Marshal(message)
}

func (j *jsonEncoder) Decode(data []byte, entity any) error {
	return json.Unmarshal(data, entity)
}
