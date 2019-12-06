package base

import (
	"fmt"
	"github.com/go-resty/resty"
	"time"
)

func RestyGet(url string, token string) ([]byte, error) {

	resp, err := resty.SetTimeout(time.Duration(60) * time.Second).R().Get(url)
	if err != nil {
		fmt.Printf("the error is %v ", err)
		return nil, err
	}

	body := resp.Body()
	//body, err := ioutil.ReadAll(response.Body)

	fmt.Printf("body=%s \n", string(body))
	return body, nil
}