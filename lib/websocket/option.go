package websocket

func WithConnPoolMaxSize(size int) Option {
	return func(server *wsServer) {
		server.connPoolMaxSize = size
	}
}
