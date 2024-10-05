/**
 * @author ysj
 * @email 2239831438@qq.com
 * @date 2024-10-05 14:10:28
 */

package ip

import (
	"net"
	"strings"
)

func GetIPv4Addr(prefix string) ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var res []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			if ipv4 := ipnet.IP.String(); strings.HasPrefix(ipv4, prefix) { // 只获取 IPv4 地址
				res = append(res, ipv4)
			}
		}
	}
	return res, nil
}
