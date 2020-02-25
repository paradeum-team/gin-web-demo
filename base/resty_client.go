package base

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty"
	"io/ioutil"
	"net"
	"time"
)

func RestyGet(url string, token string) ([]byte, error) {

	resp, err := resty.SetTimeout(time.Duration(60) * time.Second).R().Get(url)
	if err != nil {
		fmt.Printf("the error is %v ", err)
		return nil, err
	}
	defer resp.RawResponse.Body.Close()
	body := resp.Body()
	//body, err := ioutil.ReadAll(response.Body)

	fmt.Printf("body=%s \n", string(body))

	return body, nil
}


func RestyPost(url string,data interface{} ,token string) ([]byte, error) {
	resp, err := resty.SetTimeout(time.Duration(60) * time.Second).R().SetBody(data).Get(url)
	if err != nil {
		fmt.Printf("the error is %v ", err)
		return nil, err
	}
	defer resp.RawResponse.Body.Close()
	body := resp.Body()
	//body, err := ioutil.ReadAll(response.Body)

	fmt.Printf("body=%s \n", string(body))

	return body, nil
}

func RestyPostForm(filename string ,filePath string,formData map[string]string,url string  )(scode int,data []byte,err error){

	fileb, _ := ioutil.ReadFile(filePath)

	resp, err := resty.SetTimeout(time.Duration(60)*time.Second).R().SetFileReader("file", filename, bytes.NewReader(fileb)).
		SetFormData(formData).
		Post(url)

	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			fmt.Printf("connection timeout exception  ...")
		}else{
			fmt.Printf("connection exception ...http 500 ")

		}
	}

	return resp.StatusCode(),resp.Body(),err

}