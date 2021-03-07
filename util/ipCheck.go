package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)
/*构建icmp包*/
type ICMP struct {
	Type uint8 /*类型*/
	Code uint8 /*代码*/
	Checksum uint16 /*校验和*/
	Identifier uint16 /*标识符*/
	SequenceNum uint16 /*序列号*/
}
/*校验和检测*/
/*这里是校验和的计算方式*/
/*
	1.将校验和字段置为0
	2.将每两个字节（16位）相加（二进制求和）直到最后得出结果，若出现最后还剩一个字节继续与前面结果相加
	3.将高16位与低16位相加，直到高16位为0为止
	4.将最后的结果（二进制）取反
*/
func Checksum(data []byte) uint16{
	var(
		sum uint32 = 0
		length int = len(data) /*校验时返回的校验码长度*/
		index int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)
	//fmt.Println(sum)
	return uint16(^sum)
}
func ipCheck(addr string) bool {
	var(
		icmp ICMP
		laddr = net.IPAddr{IP:net.ParseIP("0.0.0.0")}
		raddr,_ = net.ResolveIPAddr("ip",addr)
	)
	conn, err := net.DialIP("ip4:icmp",&laddr,raddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer conn.Close()
	/*构建数据包内容*/
	icmp.Type = 8
	icmp.Code = 0
	icmp.Checksum = 0
	icmp.Identifier = 0
	icmp.SequenceNum = 0

	/*定义缓存空间*/
	var buffer bytes.Buffer
	/*现在buffer中写入icmp数据报获取校验和*/
	binary.Write(&buffer,binary.BigEndian,icmp)
	icmp.Checksum = Checksum(buffer.Bytes())
	//然后清空buffer，将所求完整的校验和的icmp数据报写入其中准备发送
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, icmp)

	_, err = conn.Write(buffer.Bytes())
	/*判断写入完整数据报是否成功*/
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	/*声明一个用于存储接收数据的切片*/
	recv := make([]byte, 1024)
	/*设定超时请求*/
	conn.SetReadDeadline((time.Now().Add(time.Second * 1)))
	_,err = conn.Read(recv)

	if err != nil {
		return false
	}else{
		return true
	}
}

func IpProcess(hosts []string) ([]string){

	var host_list []string
	var wg sync.WaitGroup
	for _,ip := range hosts{
		wg.Add(1)
		go func(j string){
			defer wg.Done()

			checkrs := ipCheck(j)

			if checkrs {  //如果结果为真
				fmt.Println(checkrs)
				host_list = append(host_list,j)
				fmt.Println(j,"存活!")
			}else{

				fmt.Println(j,"不存活")
			}

		}(ip)


	}
	wg.Wait()
	return host_list
}

