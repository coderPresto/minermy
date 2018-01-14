package services

import (
  //"github.com/kataras/iris"
)

type UserService interface {
  GetHello() string
}

type userService struct {

}

func (s *userService) GetHello() string {
  return "Hello World"
}
