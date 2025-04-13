// Code generated by goctl. DO NOT EDIT.
// Source: statusclient.proto

package statusclient

import (
	"context"

	"github.com/YShiJia/IM/apps/status/rpc/statusmodel"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ClientConnAddressRequest  = statusmodel.ClientConnAddressRequest
	ClientConnAddressResponse = statusmodel.ClientConnAddressResponse
	ClientMsgSyncRequest      = statusmodel.ClientMsgSyncRequest
	ClientMsgSyncResponse     = statusmodel.ClientMsgSyncResponse
	Request                   = statusmodel.Request
	Response                  = statusmodel.Response
	UserOnlineRequest         = statusmodel.UserOnlineRequest
	UserOnlineResponse        = statusmodel.UserOnlineResponse

	Status interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		UserOnline(ctx context.Context, in *UserOnlineRequest, opts ...grpc.CallOption) (*UserOnlineResponse, error)
		ClientConnAddress(ctx context.Context, in *ClientConnAddressRequest, opts ...grpc.CallOption) (*ClientConnAddressResponse, error)
		ClientMsgSync(ctx context.Context, in *ClientMsgSyncRequest, opts ...grpc.CallOption) (*ClientMsgSyncResponse, error)
	}

	defaultStatus struct {
		cli zrpc.Client
	}
)

func NewStatus(cli zrpc.Client) Status {
	return &defaultStatus{
		cli: cli,
	}
}

func (m *defaultStatus) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := statusmodel.NewStatusClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultStatus) UserOnline(ctx context.Context, in *UserOnlineRequest, opts ...grpc.CallOption) (*UserOnlineResponse, error) {
	client := statusmodel.NewStatusClient(m.cli.Conn())
	return client.UserOnline(ctx, in, opts...)
}

func (m *defaultStatus) ClientConnAddress(ctx context.Context, in *ClientConnAddressRequest, opts ...grpc.CallOption) (*ClientConnAddressResponse, error) {
	client := statusmodel.NewStatusClient(m.cli.Conn())
	return client.ClientConnAddress(ctx, in, opts...)
}

func (m *defaultStatus) ClientMsgSync(ctx context.Context, in *ClientMsgSyncRequest, opts ...grpc.CallOption) (*ClientMsgSyncResponse, error) {
	client := statusmodel.NewStatusClient(m.cli.Conn())
	return client.ClientMsgSync(ctx, in, opts...)
}
