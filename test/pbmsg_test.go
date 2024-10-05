/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-05 15:16:37
 */

package test

import (
	"fmt"
	"github.com/YShiJia/IM/lib/encoder"
	"github.com/YShiJia/IM/model/pbmessage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPbMessage(t *testing.T) {
	msg := pbmessage.PbMessage{
		MsgType: pbmessage.PbMessageType_Ping,
		From:    "testUid1",
		To:      "testUid1",
		Seq:     1,
		Err:     "",
		Data:    []byte("hello"),
	}
	jsonEncoder := encoder.NewJsonEncoder()
	data, err := jsonEncoder.Encode(&msg)
	assert.NoError(t, err)
	fmt.Println(string(data))
}
