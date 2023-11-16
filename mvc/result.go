package mvc

import (
	"fmt"
)

func Result(code int, result interface{}, msg string) ResultReq {
	return ResultReq{
		Code:   code,
		Msg:    msg,
		Result: result,
	}
}

func Ok() ResultReq {
	return Result(OK, nil, "ok")
}
func OkWithMsg(msg string) ResultReq {
	return Result(OK, nil, msg)
}

func Success(result interface{}) ResultReq {
	return Result(OK, result, "ok")
}

func Error(error string) ResultReq {
	return ResultReq{
		Code: Err,
		Msg:  error,
	}
}

func Fail(code int, msg string) ResultReq {
	return Result(code, nil, msg)
}

func FailWithMsg(code int, msg string) ResultReq {
	return Result(code, nil, msg)
}

func HandleError(err error, codes ...interface{}) {
	if err != nil {
		if len(codes) > 0 {
			panic(fmt.Sprintf("%v\n%s", codes, err.Error()))
		} else {
			panic(err.Error())
		}
	}
}

type ResultReq struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}

type Page struct {
	PageCount int         `json:"pageCount"`
	PageNo    int         `json:"pageNo"`
	PageSize  int         `json:"pageSize"`
	Total     int         `json:"total"`
	Items     interface{} `json:"items"`
}
