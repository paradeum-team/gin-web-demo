package utils

import (
	"fmt"
	"os/exec"
)

func Open(){
	f, err := exec.Command("open", "/Users/xcdong/gitHisun/gin-web-demo/swagger.md").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("******* \n")
	fmt.Println(string(f))
	fmt.Println("123")
}


