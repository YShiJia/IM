/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-01 18:52:04
 */

package webscoket

type HandleFunc func(ctx *WebContext)

type Middleware = HandleFunc

type Route interface {
	Register(method, path string, handleFunc HandleFunc) bool
	Get(method, path string) (HandleFunc, bool)
}

var _ Route = (*routeByMap)(nil)

// TODO 现在先暂时使用map实现，后续开源一个前缀树节点项目，使用其实现路由树内核逻辑
type routeByMap struct {
	//普通map并发读，没问题
	routes map[string]map[string]HandleFunc
}

func NewRouteByMap() *routeByMap {
	return &routeByMap{
		routes: make(map[string]map[string]HandleFunc),
	}
}

// Register 若是注册已存在路由,返回false
// 不可并发写
func (r *routeByMap) Register(method, path string, handleFunc HandleFunc) bool {

	_, ok := r.routes[method]
	if !ok {
		r.routes[method] = make(map[string]HandleFunc)
	}
	_, ok = r.routes[method][path]
	if ok {
		//节点已存在，返回添加失败false
		return false
	}
	r.routes[method][path] = handleFunc
	return true
}

// Get 获取路由
func (r *routeByMap) Get(method, path string) (HandleFunc, bool) {
	_, ok := r.routes[method]
	if !ok {
		return nil, false
	}
	f, ok := r.routes[method][path]
	if !ok {
		return nil, false
	}
	return f, true
}
