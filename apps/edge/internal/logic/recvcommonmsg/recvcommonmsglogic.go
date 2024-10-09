/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-03 12:55:56
 */

package recvcommonmsg

import (
	"context"
	"fmt"
	"github.com/YShiJia/IM/apps/edge/internal/svc"
	"github.com/YShiJia/IM/pbmodel/pbmessage"
	"github.com/zeromicro/go-zero/core/logx"
)

type RecvCommonMsgConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecvCommonMsgConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *RecvCommonMsgConsumer {
	return &RecvCommonMsgConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (r *RecvCommonMsgConsumer) Consume(ctx context.Context, key, value string) error {
	if key != r.svcCtx.Name {
		//判断不是发送给当前服务的cmd消息
		logx.Errorf("[RecvCmdMsgConsumer] cmd msg is not belong to local server, key: %s", key)
		return fmt.Errorf("cmd msg is not belong to local server, key: %s", key)
	}
	//对消息进行解码
	var msg pbmessage.PbMessage
	if err := r.svcCtx.Encoder.Decode([]byte(value), &msg); err != nil {
		logx.Errorf("[RecvCmdMsgConsumer] proto.Unmarshal failed, err: %v", err)
		return err
	}
	return r.handleCmdMsg(&msg)
}

func (r *RecvCommonMsgConsumer) handleCmdMsg(msg *pbmessage.PbMessage) error {
	//使用策略模式进行填充
	return CommonMsgProcessor.CommonMsgHandle(r, msg)
	return nil
}
