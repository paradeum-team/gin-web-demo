package base

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"gin-web-demo/common/dict"
	"gin-web-demo/common/utils"
	pldconf "gin-web-demo/config"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)
func HttpGet(url string,token string ) ([] byte, error) {
	method:=dict.HttpGet
	return httpExecute4JSON(method,url,"",token)
}

func HttpDelete(url string,token string ) ([] byte, error) {
	method:=dict.HttpDelete
	return httpExecute4JSON(method,url,"",token)
}

func HttpPut(url string,jsonBody string,token string ) ([] byte, error) {
	method:=dict.HttpPut
	return httpExecute4JSON(method,url,jsonBody,token)
}

func HttpPost(url string,jsonBody string,token string ) ([] byte, error) {
	method:=dict.HttpPost
	return httpExecute4JSON(method,url,jsonBody,token)
}

func HttpForm(url string,params map[string]interface{}, filePath string, token string) ([] byte, error){
	method:=dict.HttpPost
	return httpExecute4PostFile(method,url,params,filePath,token)
}
/**
 * 获取 http client
 * 忽略证书
 *
 */
func getHttpClient() *http.Client {
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

	if pldconf.AppConfig.Server.ProxyModel { //本地开发模式，需要使用代理
		tr.Proxy = http.ProxyURL(proxy)
	}
	client := &http.Client{
		Transport: tr,
		//Timeout:   time.Second * 5, //超时时间
	}
	return client
}

/**
 * jsonParams:`{"userName":"DongXC"}`
 */
func httpExecute4JSON(httpMethod string, url string, jsonParams string, auth string) ([] byte, error) {
	var requestBody *strings.Reader //默认值 nil
	if   len(jsonParams) > 0 {
		requestBody = strings.NewReader(jsonParams)
	}

	return httpExcute(httpMethod,url,requestBody,auth)
}
/**
* http client get request
* return httpStatusCode ,body ,err
*/
func httpExecute4PostFile(httpMethod string,url string,params map[string]interface{}, filePath string, auth string) ([] byte, error) {
	var message = ""
	file, err := os.Open(filePath)
	if err != nil {
		message=fmt.Sprintf("[HttpExecute method=POST]  open file is  error. %v \n",err)
		log.Println(message)
		return nil, errors.New(message)
	}
	defer file.Close()

	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	part, err := writer.CreateFormFile("file", getTempPath()+file.Name())
	if err != nil {
		message=fmt.Sprintf("[HttpExecute method=POST]  CreateFormFile is  wrong. %v \n",err)
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		switch v := val.(type) {
		case string:
			_ = writer.WriteField(key, v)
		case int:
			var s int
			s = v
			v2 := strconv.Itoa(s)
			_ = writer.WriteField(key, v2)
		case bool:
			var s bool
			s = v
			v2 := strconv.FormatBool(s)
			_ = writer.WriteField(key, v2)
		default:
			message=fmt.Sprintf("convert type error.the param type  must be in [string,int ,bool]")
			log.Println(message)
			return nil,errors.New(message)
		}

	}
	err = writer.Close()
	if err != nil {
		message=fmt.Sprintf("[HttpPostExecute]  close the temp writer is  wrong. %v \n",err)
		return nil, err
	}

	return httpExcute(httpMethod,url,requestBody,auth)

}


func httpExcute(httpMethod string, url string, requestBody io.Reader, auth string) ([] byte, error) {
	client := getHttpClient()
	var message = ""
	req, err := http.NewRequest(httpMethod, url, requestBody)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	if auth != "" {
		req.Header.Set("Authorization", auth)
	}

	response, err := client.Do(req)

	if err != nil {
		message=fmt.Sprintf("请求错误：%v", err)
		log.Println(message)
		return nil, errors.New(message)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		message=fmt.Sprintf("读取body失败：%v", err)
		log.Println(message)
		return nil, errors.New(message)
	}

	return   body, nil
}
func getTempPath()string {

	timeContent:=utils.GetCurrentTimeUnix()
	return "./data/temp/" +timeContent

}
