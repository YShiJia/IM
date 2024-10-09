/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-05 14:07:04
 */

package test

import (
	"fmt"
	"github.com/YShiJia/IM/lib/ip"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestGetIPV4(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // 只获取 IPv4 地址
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
func TestGetIPV4WithPrefix(t *testing.T) {
	prefix := "172"
	addr, err := ip.GetIPv4Addr(prefix)
	assert.NoError(t, err)
	for _, ipv4 := range addr {
		fmt.Println(ipv4)
	}
}
