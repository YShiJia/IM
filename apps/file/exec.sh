

goctl api go -api ./apps/file/file.api -dir ./apps/file -style go_zero

goctl api doc -dir ./apps/message/api -o ./apps/message/api

goctl rpc protoc ./apps/message/rpc/message.proto --go_out=./apps/message/rpc/ --go-grpc_out=./apps/message/rpc/ --zrpc_out=./apps/message/rpc/
