package main

import (
	"log"
	"net/url"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	proxyAddr := "http://121.232.148.165:9000"
	httpUrl := "http://www.baidu.com"
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	res, err := httpClient.Get(httpUrl)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}
	c, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(c))
}
