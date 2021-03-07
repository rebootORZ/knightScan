package util

import (
	"fmt"
	"github.com/badoux/goscraper"
	"strings"
	"sync"
)
func HttpBanner(hosts []string) { //暂时没有banners返回，后续完善指纹识别的时候会用到
	//var banners = []string{}
	var wg sync.WaitGroup
	for _,v:=range hosts{
		wg.Add(1)
		go func(j string){
			defer wg.Done()
			url := "http://" + j

			rs, err := goscraper.Scrape(url, 5)
			if err != nil {
				fmt.Println(url + "    【访问出错，请手动尝试】")
			}else{
				if rs != nil{
					httpsflag := strings.Contains(rs.Preview.Title, "HTTPS") //判断标题是否为：400 The plain HTTP request was sent to HTTPS port
					if httpsflag{
						url2 := "https://" +j
						rs2, err2 := goscraper.Scrape(url2, 5)
						if err2 != nil {
							fmt.Println(url2 + "    【访问出错，请手动尝试】")

						}else {
							if rs2 != nil{
								fmt.Printf("%s  %s\n", url2,rs2.Preview.Title)
							}

						}

					}else{
						fmt.Printf("%s : %s\n", url,rs.Preview.Title)
					}

				}
			}

		}(v)

	}
	wg.Wait()
	//return banners
}
