package example

import (
	"context"
	"github.com/booldesign/codes"
	"github.com/pkg/errors"
	"testing"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2022/4/9 19:37
 * @Desc:
 */

const (
	ErrSuccess int = iota + 10000
	ErrUnknown
	ErrRecordNotFind
	ErrValidate
)

var (
	ErrRecordNotFindErr = errors.New("record not find")
)

type User struct {
	UserId   int64  `gorm:"Column:user_id"`  // 用户id
	Avatar   string `gorm:"Column:avatar"`   // 头像
	Nickname string `gorm:"Column:nickname"` // 昵称
}

func init() {
	register(ErrSuccess, 200, "OK")
	register(ErrUnknown, 500, "Internal server error")
	register(ErrRecordNotFind, 404, "Record not find")
	register(ErrValidate, 404, "Params error")
}

func TestWithCode(t *testing.T) {
	postHandler(User{
		UserId:   1,
		Nickname: "hello",
	})
}

func postHandler(user User) {
	err := get(user)
	if err != nil {
		// 添加一些上下文信息包装次错误
		err = codes.WrapCode(err, ErrValidate, "user id:%d err", user.UserId)
		Response(context.TODO(), err, nil)
		return
	}
	Response(context.TODO(), nil, "SUCCESS")
}

func get(user User) error {
	err := dbQuery(user)
	if err != nil {
		return err
	}
	return nil
}

func dbQuery(user User) error {
	err := ErrRecordNotFindErr
	return codes.WithCode(ErrRecordNotFind, err.Error())
}
