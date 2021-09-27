package web

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"stk/ecode"
	"stk/xlog"
)

type Data struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Json(c *gin.Context, data interface{}) {
	if e, ok := data.(ecode.Codes); ok {
		req, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			xlog.Errorf("ioutil.ReadAll(%v),err:%v", c.Request.Body, err)
		}
		xlog.Errorf("api error-->url:%s,params:%s,err:%v", c.Request.URL.Host+c.Request.URL.Path+c.Request.URL.RequestURI(), string(req), e)
		c.JSON(200, Data{
			Code: e.Code(),
			Msg:  e.Message(),
			Data: nil,
		})
	} else {
		c.JSON(
			200,
			Data{
				Code: 1,
				Msg:  "success",
				Data: data,
			},
		)
	}
}
