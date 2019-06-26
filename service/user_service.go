package service

import (
	"fmt"
	"gin-web-demo/common/utils"
	"gin-web-demo/entity"
)

type UserService interface {
	GetInfo(name string) entity.User
	ListUsers() []entity.User
}
//定义一个全局的私有变量，存放service 的实例。避免每次都生一个实例，减少垃圾回收的成本。
var userServiceInstace *userService
type userService struct{

}



func NewUserServiceInstace() UserService{
	if userServiceInstace ==nil{
		userServiceInstace=&userService{}
	}
	return userServiceInstace
}

func (*userService) GetInfo(name string) entity.User {
	plogger.NewInstance().GetLogger().Info("Call the method , GetInfo ...logger.level=info ")
	plogger.NewInstance().GetLogger().Debug("Call the method , GetInfo ...logger.level=debug  ")
	fmt.Println("you will get userinfo by name=",name)

	return entity.User{Name:"dxc",Age:30,Code:"203462",Address:"回龙观"}
}

func (*userService) ListUsers() []entity.User {
	panic("implement me")

	user1 :=entity.User{Name:"dxc",Age:30,Code:"203462",Address:"回龙观"}
	user2:=entity.User{Name:"dxc",Age:30,Code:"203462",Address:"回龙观"}
	users :=[]entity.User{user1,user2}
	return users


}