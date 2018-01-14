package controllers

import (
  "minermy/services"

  //"github.com/kataras/iris"
)

type UserController struct {
  Service services.UserService
}

func (c *UserController) GetTest() string {
  testvar := c.Service.GetHello();
  return testvar
  //return c.Service.GetHello();
}
