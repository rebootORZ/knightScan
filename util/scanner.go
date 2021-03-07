package util

import (
	"net"
	"time"
	"sync"
)

var ports = [...]string{"80","8080","443","7001","9090","81"}  //常见web端口  ,"80","8080","443","7001","9090","81"
var ip_ports = []string{} //用来存放开放端口的 IP-端口 信息

func Scanner(hosts []string) []string{
	var wg sync.WaitGroup

	for _, ip := range hosts{
		wg.Add(1)
		go func(j string){
			defer wg.Done()
			for _, port := range ports{
				address := j + ":" + port
				d := net.Dialer{Timeout: time.Second*1} //这里得设置下延迟
				conn, err := d.Dial("tcp", address)
				if err != nil {
					continue
				}
				conn.Close()
				ip_ports = append(ip_ports,address)
			}

		}(ip)

	}
	wg.Wait()
	return ip_ports
}