// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"github.com/YShiJia/IM/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	EmailVerifyCodeReq  = user.EmailVerifyCodeReq
	EmailVerifyCodeResp = user.EmailVerifyCodeResp
	LoginReq            = user.LoginReq
	LoginResp           = user.LoginResp
	RegisterReq         = user.RegisterReq
	RegisterResp        = user.RegisterResp
	Request             = user.Request
	Response            = user.Response
	UserEntity          = user.UserEntity

	User interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		EmailVerifyCode(ctx context.Context, in *EmailVerifyCodeReq, opts ...grpc.CallOption) (*EmailVerifyCodeResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) EmailVerifyCode(ctx context.Context, in *EmailVerifyCodeReq, opts ...grpc.CallOption) (*EmailVerifyCodeResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EmailVerifyCode(ctx, in, opts...)
}
