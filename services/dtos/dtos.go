package dtos

import "myblogs/util/err"

type LoginDto struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type LoginResultDto struct {
	Name  string `json:"name"`
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

// ResultData return obj
type ResultData struct {
	Code  err.ErrCode `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Total int         `json:"total,omitempty"`
	Msg   string      `json:"msg,omitempty"`
}

// ErrMsg 获取错误信息
func ErrMsg(code err.ErrCode, msg ...interface{}) string {
	if msg != nil {
		return msg[0].(string)
	}
	return err.ErrMsg[code]
}

// Ok 参数objs[0]=object,objs[1]=string
func Ok(objs ...interface{}) *ResultData {
	data := &ResultData{
		Code: err.Success,
	}
	if objs == nil {
		data.Msg = ErrMsg(err.Success)
		return data
	}

	if len(objs) == 1 {
		data.Data = objs[0]
		return data
	}
	return &ResultData{
		Code: err.Success,
		Data: objs[0],
		Msg:  objs[1].(string),
	}
}

// NotOk 参数objs[0]=string,objs[1]=object
func NotOk(code err.ErrCode, objs ...interface{}) *ResultData {
	data := &ResultData{
		Code: code,
	}
	if objs == nil {
		data.Msg = ErrMsg(code)
		return data
	}
	if len(objs) == 1 {
		data.Msg = objs[0].(string)
		return data
	}
	return &ResultData{
		Code: err.Success,
		Data: objs[1],
		Msg:  objs[0].(string),
	}
}
