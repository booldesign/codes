package example

import (
	"context"
	"fmt"
	"github.com/booldesign/codes"
	"log"
	"net/http"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/4/10 11:51
 * @Desc:
 */

type ErrResponse struct {
	Code int `json:"code"`

	Message string `json:"message"`

	Data interface{}
}

func Response(c context.Context, err error, data interface{}) {
	if err != nil {
		log.Printf("%#+v\n", err)
		coder := codes.ParseCoder(err)
		// write json
		fmt.Printf("http code %d, data:%+v\n", coder.HTTPStatus(), ErrResponse{
			Code:    coder.Code(),
			Message: coder.String(),
		})
		return
	}
	// write json
	fmt.Printf("http code %d, data:%+v\n", http.StatusOK, data)
}
