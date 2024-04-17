package main

import (
	"fmt"
	"net"
)


func getIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("无法获取网络接口：", err)
		return	
	}
	var result []string
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // 当前计算机的IPv4地址：
				result = append(result, ipnet.IP.String())
			}
		}
	}
	fmt.Println("getIp result", result)
}

func getIpV2() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("无法获取网络接口：", err)
		return
	}

	for _, iface := range interfaces {
		addresses, err := iface.Addrs()
		if err != nil {
			fmt.Println("无法获取接口地址：", err)
			continue
		}

		for _, addr := range addresses {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					fmt.Println("当前计算机的IPv4地址：", ipNet.IP.String())
				} else if ipNet.IP.To16() != nil {
					// fmt.Println("当前计算机的IPv6地址：", ipNet.IP.String())
				}
			}
		}
	}
}

func main() {
	getIp()
	getIpV2()
}