package router

import (
	"log"
	"myblogs/dtos"
	"myblogs/middleware"
	"myblogs/util/err"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Login 登录
func Login(c *gin.Context) {
	loginDto := &dtos.LoginDto{}

	e := c.ShouldBindJSON(loginDto)
	log.Println(loginDto)
	if e != nil {
		log.Println(e)
		c.JSON(http.StatusOK, &dtos.ResultData{
			Code: err.CertErr,
			Msg:  dtos.ErrMsg(err.CertErr, "用户名/密码错误"),
		})
		return
	}
	name := viper.GetString("security.name")
	pwd := viper.GetString("security.pwd")
	if name == loginDto.Name && pwd == loginDto.Pwd {

		token, exp := middleware.CreateJwtToken(name)
		c.JSON(http.StatusOK, dtos.Ok(&dtos.LoginResultDto{
			Name:  loginDto.Name,
			Token: token,
			Exp:   exp,
		}))
		return

	}
	log.Println("登录失败")
	c.JSON(http.StatusOK, dtos.NotOk(err.CertErr))
}
