/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2025-05-04 15:27:33
 */

package message

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	conf "github.com/YShiJia/IM/apps/message/api/internal/config"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/db"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/mq"
	"github.com/YShiJia/IM/apps/message/api/internal/dao/redisdb"
	"github.com/YShiJia/IM/model"
	"github.com/YShiJia/IM/model/ext"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Transfer() {
	for {
		message, err := mq.SendMsgQueueReader.ReadMessage(context.TODO())
		if err != nil {
			log.Errorf("read message from kafka error: %v", err)
			continue
		}
		msg := &ext.Message{}
		if err := json.Unmarshal(message.Value, msg); err != nil {
			log.Errorf("message key[%s] recv data %s, unmarshal error: %v", message.Key, string(message.Value), err)
		}
		switch msg.Type {
		case model.MessageTypePrivate:
			if err := transferPrivateMessage(message.Value, msg); err != nil {
				log.Errorf("message key[%s] recv data %s, transfer error: %v", message.Key, string(message.Value), err)
			}
		case model.MessageTypeGroup:
			if err := transferGroupMessage(message.Value, msg); err != nil {
				log.Errorf("message key[%s] recv data %s, transfer error: %v", message.Key, string(message.Value), err)
			}
		default:
			log.Errorf("message key[%s] recv data %s, type invalid", message.Key, string(message.Value))
		}
	}
}

// 1. 查询接收者
// 2. 解析消息，落库
// 3. 查询对应的edge节点的RecvMQ
// 4. 发送消息到对应的edge节点的RecvMQ
func transferPrivateMessage(data []byte, msg *ext.Message) error {
	_, err := db.User.GetByUID(msg.To)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("接收者用户[UID:%s]不存在", msg.To)
		}
		log.Errorf("用户[UID:%s]查询错误 err %v", msg.To, err)
	}

	content := datatypes.JSONMap{}
	if err = content.UnmarshalJSON(msg.Content); err != nil {
		log.Errorf("解析消息内容失败 err %v", err)
		return err
	}
	pg := &model.PrivateMessage{
		FromUID:  msg.From,
		ToUID:    msg.To,
		UnionUID: genernalUnionUID(msg.From, msg.To),
		SendTime: msg.SendTime,
		Type:     uint8(msg.ContentType),
		Content:  content,
	}

	// 消息落库
	if err := db.MessagePrivate.Create(pg); err != nil {
		log.Errorf("用户[UID:%s]保存错误 err %v", msg.To, err)
		return err
	}

	// 查询对应的edge节点的RecvMQ
	edgeName, err := redisdb.GetValue(context.TODO(), fmt.Sprintf("%s-%s", conf.Conf.RedisUserInfoPrefix, msg.To))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// 用户不在线
			return nil
		}
		log.Errorf("查询用户[UID:%s]对应的edge节点失败 err %v", msg.To, err)
		return err
	}

	if writer := mq.GetRecvMsgQueueWriter(edgeName); writer != nil {
		// 落库
		if err := writer.WriteMessages(context.TODO(), kafka.Message{
			Key:   []byte(msg.To),
			Value: data,
		}); err != nil {
			log.Errorf("写入消息失败 err %v", err)
			return err
		}
	}
	return nil
}

// 1. 查询接收群
// 2. 落库
// 3. 查询对应的edge节点的RecvMQ
// 4. 发送消息到对应的edge节点的RecvMQ
func transferGroupMessage(data []byte, msg *ext.Message) error {
	group, err := db.Group.GetByUID(msg.To, db.Group.PreloadGroupMembers())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("接收群[UID:%s]不存在", msg.To)
		}
		log.Errorf("群聊[UID:%s]查询错误 err %v", msg.To, err)
	}

	content := datatypes.JSONMap{}
	if err = content.UnmarshalJSON(msg.Content); err != nil {
		log.Errorf("解析消息内容失败 err %v", err)
		return err
	}
	gg := &model.GroupMessage{
		FromUID:  msg.From,
		GroupUID: msg.To,
		SendTime: msg.SendTime,
		Type:     uint8(msg.ContentType),
		Content:  content,
	}

	// 消息落库
	if err := db.MessageGroup.Create(gg); err != nil {
		log.Errorf("用户[UID:%s]保存错误 err %v", msg.To, err)
		return err
	}

	for _, gm := range group.GroupMembers {
		// 查询对应的edge节点的RecvMQ
		edgeName, err := redisdb.GetValue(context.TODO(), fmt.Sprintf("%s-%s", conf.Conf.RedisUserInfoPrefix, gm.User.UID))
		if err != nil {
			if errors.Is(err, redis.Nil) {
				// 用户不在线
				return nil
			}
			log.Errorf("查询用户[UID:%s]对应的edge节点失败 err %v", msg.To, err)
			return err
		}
		if writer := mq.GetRecvMsgQueueWriter(edgeName); writer != nil {
			// 落库
			if err := writer.WriteMessages(context.TODO(), kafka.Message{
				Key:   []byte(msg.To),
				Value: data,
			}); err != nil {
				log.Errorf("写入消息失败 err %v", err)
				return err
			}
		}
	}
	return nil
}

func genernalUnionUID(from, to string) string {
	if from < to {
		return from + "-" + to
	}
	return to + "-" + from
}
