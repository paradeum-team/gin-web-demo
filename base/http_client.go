package base

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	pldconf "gin-web-demo/config"
)

/**
 * 获取 http client
 * 忽略证书
 *
 */
func GetHttpClient() *http.Client {
	proxyUrl := "socks5://127.0.0.1:10081"
	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		//Proxy: http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
	}

	if pldconf.AppConfig.Server.ProxyModel{//本地开发模式，需要使用代理
		tr.Proxy=http.ProxyURL(proxy)
	}
	client := &http.Client{
		Transport: tr,
		//Timeout:   time.Second * 5, //超时时间
	}
	return client
}



func httpBaseQueryExecute(url string) ([] byte, error) {
	client := GetHttpClient()

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")

	if err != nil {
		log.Fatalln("New Request is wrong :", err)
		return nil, err
	}

	userName:="admin"
	userPassword:="admin123"

	//设置用户名密码访问
	req.SetBasicAuth(userName, userPassword)

	response, err := client.Do(req)

	if err != nil { //http connection error
	//todo 可以增加 server code ，在处理的时候，知道是服务器 connection exception
		log.Println("请求错误：", err)
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("读取body失败：", err)
		return nil, err
	}

	return body, nil

}