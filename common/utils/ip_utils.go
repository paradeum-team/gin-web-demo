package utils

import (
	"log"
	"net"
	"regexp"
	"strings"
)

/**
 * 获取本地的ip <br/>
 * return [] string
 * 可能有双网卡，或者多个ip 。数组最后一个是默认的ip：127.0.0.1
 */
func GetLocalHostIps() (ipHosts [] string) {
	address, err := net.InterfaceAddrs()
	localhost := "127.0.0.1"
	ips := make([]string, 0)
	if err != nil {
		log.Printf("Error: get lochost ip is wrong ..")
		ips = append(ips, localhost)
		return ips
	}

	for _, addr := range address {
		IpDr := addr.String()
		match, _ := regexp.MatchString(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+/[0-9]+$`, IpDr)
		if !match {
			continue
		}
		ip := strings.Split(IpDr, "/")[0]
		//fmt.Printf("ip[%d]=%s \n",idx,ip)
		if localhost != ip {
			ips = append(ips, ip)
		}
	}
	ips = append(ips, localhost)

	return ips
}
