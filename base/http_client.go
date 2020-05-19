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
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func HttpGet(url string, token string, customHeader map[string]string) ([] byte, error) {
	method := dict.HttpGet
	return httpExecute4JSON(method, url, "", token, customHeader)
}

func HttpDelete(url string, token string, customHeader map[string]string) ([] byte, error) {
	method := dict.HttpDelete
	return httpExecute4JSON(method, url, "", token, customHeader)
}

func HttpPut(url string, jsonBody string, token string, customHeader map[string]string) ([] byte, error) {
	method := dict.HttpPut
	return httpExecute4JSON(method, url, jsonBody, token, customHeader)
}

func HttpPost(url string, jsonBody string, token string, customHeader map[string]string) ([] byte, error) {
	method := dict.HttpPost
	return httpExecute4JSON(method, url, jsonBody, token, customHeader)
}

func HttpForm(url string, params map[string]string, filePath string, token string, customHeader map[string]string) ([] byte, error) {
	method := dict.HttpPost
	return httpExecute4PostFile(method, url, params, filePath, token, customHeader)
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
		DisableKeepAlives:   false, // 关注 http 的连接释放；
	}

	if pldconf.AppConfig.Server.ProxyModel { //本地开发模式，需要使用代理
		tr.Proxy = http.ProxyURL(proxy)
	}
	timeout := 60
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * time.Duration(timeout), //超时时间
	}
	return client
}

/**
 * jsonParams:`{"userName":"DongXC"}`
 */
func httpExecute4JSON(httpMethod string, url string, jsonParams string, auth string, customHeader map[string]string) ([] byte, error) {
	var requestBody *strings.Reader //默认值 nil
	if len(jsonParams) > 0 {
		requestBody = strings.NewReader(jsonParams)
	}

	return httpExe4JSON(httpMethod, url, requestBody, auth, customHeader)
}

/**
* http client get request
* return httpStatusCode ,body ,err
*/
func httpExecute4PostFile(httpMethod string, url string, params map[string]string, filePath string, auth string, customHeader map[string]string) ([] byte, error) {
	var message = ""
	file, err := os.Open(filePath)
	if err != nil {
		message = fmt.Sprintf("[HttpExecute method=POST]  open file is  error. %v \n", err)
		log.Println(message)
		return nil, errors.New(message)
	}
	defer file.Close()

	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	part, err := writer.CreateFormFile("file", getTempPath()+file.Name())
	if err != nil {
		message = fmt.Sprintf("[HttpExecute method=POST]  CreateFormFile is  wrong. %v \n", err)
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		message = fmt.Sprintf("[HttpPostExecute]  close the temp writer is  wrong. %v \n", err)
		return nil, err
	}
	contentType := writer.FormDataContentType()
	return httpExecute4Form(httpMethod, url, contentType, requestBody, auth, customHeader)

}

func httpExe4JSON(httpMethod string, url string, requestBody io.Reader, auth string, customHeader map[string]string) ([] byte, error) {
	contentType := "application/json"
	//如果还有自定义的header，需要进一步透传出来。
	return httpExcute(httpMethod, url, contentType, requestBody, auth, customHeader)
}

func httpExecute4Form(httpMethod string, url string, contentType string, requestBody io.Reader, auth string, customHeader map[string]string) ([] byte, error) {
	//如果还有自定义的header，需要进一步透传出来。
	return httpExcute(httpMethod, url, contentType, requestBody, auth, customHeader)
}

/**
 * 此公共方法，在简单响应数据的时候，返回切片是可以的。但是大并发或者下载响应比较大的时候，这时，返回切片是存在问题。需要按照接口"io.reader"返回响应流。
 */
func httpExcute(httpMethod string, url string, contentType string, requestBody io.Reader, auth string, customHeader map[string]string) ([] byte, error) {
	var message = ""
	var err error

	httpStatusCode, responseBodyStream, err := httpExcuteWithReader(httpMethod, url, contentType, requestBody, auth, customHeader)
	if responseBodyStream != nil {
		defer responseBodyStream.Close()
	}

	if httpStatusCode != http.StatusOK {
		message = fmt.Sprintf("服务器响应错误：%v", err)
		log.Println(message)
		return nil, errors.New(message)
	}

	body, err := ioutil.ReadAll(responseBodyStream) //每次解析的时候，都需重新开辟： bytes.buffer
	if err != nil {
		message = fmt.Sprintf("读取body失败：%v", err)
		log.Println(message)
		return nil, errors.New(message)
	}

	return body, nil
}

func httpExcuteWithReader(httpMethod string, url string, contentType string, requestBody io.Reader, auth string, customHeader map[string]string) (int, io.ReadCloser, error) {
	client := getHttpClient()
	var message = ""
	var req *http.Request
	var err error

	if dict.HttpGet == httpMethod {
		req, err = http.NewRequest(httpMethod, url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	} else {
		req, err = http.NewRequest(httpMethod, url, requestBody)
		req.Header.Set("Content-Type", contentType)

	}
	//req.Header.Set("Connection", "keep-alive")

	if auth != "" {
		req.Header.Set("Authorization", auth)
	}

	for key, value := range customHeader {
		req.Header.Set(key, value)
	}

	response, err := client.Do(req)
	if err != nil {
		message = fmt.Sprintf("请求错误：%v", err)
		log.Println(message)

		return -1, nil, errors.New(message)
	}

	/**
	    // drain body and close
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	 */

	return response.StatusCode, response.Body, nil
}
func getTempPath() string {
	//rand.Seed(time.Now().UnixNano())
	timeContent := utils.GetCurrentTimeUnix()
	randnum := strconv.Itoa(rand.Intn(9223372036854775805))
	filename := timeContent + randnum
	return "./data/temp/" + filename

}
