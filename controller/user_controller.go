package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"login-service/model"
	"login-service/service"
	"net/http"
)

type internalError struct {
	Message string `json:"message"`
}

// 707324938877.dkr.ecr.us-east-1.amazonaws.com/amakosi.dev

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	v1 := r.Group("/api/v1/")
	{
		v1.POST("auth", createLogin)
		v1.GET("login/:user_name/:password", login)
	}
	return r
}

func createLogin(c *gin.Context) {
	p := model.User{}
	err := c.Bind(&p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	err = service.AddUser(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, p)
}

func login(c *gin.Context) {
	userName := c.Param("user_name")
	passWord := c.Param("password")
	if len(userName) == 0 && len(passWord) == 0 {
		c.JSON(http.StatusInternalServerError, generateErrorMessage("username/password combination required"))
		return
	}
	res, err := service.Login(userName, passWord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, generateErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, res)
}

func generateErrorMessage(e string) string {
	ie := internalError{Message: e}
	buf, err := json.Marshal(ie)
	if err != nil {
		return fmt.Sprintf(`{"message": "%s"}`, e)
	}
	return string(buf)
}
