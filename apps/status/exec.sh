goctl rpc protoc ./apps/status/rpc/status.proto --go_out=./apps/status/rpc/ --go-grpc_out=./apps/status/rpc/ --zrpc_out=./apps/status/rpc/

goctl api go -api ./apps/status/api/status.api -dir ./apps/status/api/ -style gozero