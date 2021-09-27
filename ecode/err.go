package ecode

import (
	"sync"
)

var (
	ErrorMap = new(sync.Map)
	// base error
	OK         = New(1, "success")
	RequestErr = New(20000, "Parameter error")
	UserIDErr  = New(30000, "用户身份验证失败。")
	DevIDErr   = New(30001, "开发者身份验证失败。")
	SignErr    = New(40000, "签名验证失败。")
	ServerErr  = New(50000, "服务异常。")
	GatewayErr = New(50001, "网关异常。")
	UserExist  = New(60001, "user exist")
)
