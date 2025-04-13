/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 18:52:04
 */

package websocket

type Route interface {
	Register(method, path string, handleFunc HandleFunc)
	Get(method, path string) (HandleFunc, bool)
}

var _ Route = (*routeByMap)(nil)

// TODO 现在先暂时使用map实现，后续开源一个前缀树节点项目，使用其实现路由树内核逻辑
type routeByMap struct {
	//普通map并发读，没问题
	routes map[string]map[string]HandleFunc
}

func NewRouteByMap() Route {
	return &routeByMap{
		routes: make(map[string]map[string]HandleFunc),
	}
}

// Register 重复注册会覆盖前面的逻辑，不可并发写
func (r *routeByMap) Register(method, path string, handler HandleFunc) {
	_, ok := r.routes[method]
	if !ok {
		r.routes[method] = make(map[string]HandleFunc)
	}
	r.routes[method][path] = handler
}

// Get 获取路由
func (r *routeByMap) Get(method, path string) (HandleFunc, bool) {
	methodMap, ok := r.routes[method]
	if !ok {
		return nil, false
	}

	handler, ok := methodMap[path]
	if !ok {
		return nil, false
	}

	return handler, true
}
