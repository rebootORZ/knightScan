package main

import (
	"./util"
	"flag"
	"fmt"
)

func main() {
	//util.HttpBanner()
	fmt.Println("——————————————————————————————————————")
	fmt.Println("|####################################|")
	fmt.Println("|       Powered by RebootORZ         |")
	fmt.Println("|         Date 2021.03.07            |")
	fmt.Println("|####################################|")
	fmt.Println("——————————————————————————————————————")

	var cidr string
	flag.StringVar(&cidr,"h","","IP范围")
	flag.Parse()  //转换
	hosts,_ := util.GetIps(cidr)  //传入用户输入的IP段

	minIp := hosts[0]
	maxIp := hosts[len(hosts)-1]
	fmt.Println("探测范围：" + minIp + "~" + maxIp)


	ipHosts := util.Scanner(hosts)
	util.HttpBanner(ipHosts)
	//fmt.Println(titles)



}