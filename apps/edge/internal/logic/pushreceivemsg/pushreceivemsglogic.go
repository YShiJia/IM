/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-02 18:35:57
 */

package pushreceivemsg

import (
	"context"
	"github.com/YShiJia/IM/apps/edge/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PushReceiveMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushReceiveMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushReceiveMsgLogic {
	return &PushReceiveMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PushReceiveMsgLogic 将传入的msg序列化为str， 推送至对应的MQ
// 2024年10月2日18:46:07，这里有一个坑，KPush可能会造成ctx数据堆积，得不到释放，所以每一次调用这个函数，需要使用新的ctx
// TODO 后期再来想一个更好的解决方案
func (l *PushReceiveMsgLogic) PushReceiveMsgLogic(key string, msg any) error {
	data, err := l.svcCtx.Encoder.Encode(msg)
	if err != nil {
		return err
	}
	return l.svcCtx.SendMsgPusher.KPush(l.ctx, key, string(data))
}
