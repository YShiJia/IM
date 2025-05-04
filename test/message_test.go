/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-04 17:11:57
 */

package test

import (
	"encoding/json"
	"github.com/YShiJia/IM/model"
	"github.com/YShiJia/IM/model/ext"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_Message(t *testing.T) {
	msg := ext.Message{
		Type:     model.MessageTypePing,
		From:     "uid-u1",
		To:       "service",
		SendTime: time.Now().Unix(),
		Content:  []byte("hello"),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	log.Infof(string(data))
	m := ext.Message{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		t.Error(err)
	}
}
