package errorx

import (
	"github.com/pkg/errors"
)

// 定义mysql错误类型
const (
	MYSQL_FIND_ERR = iota + 1000
	MYSQL_COUNT_ERR
)

// 定义业务错误类型
const (
	HTTP_UNKNOW_ERR = iota + 2000
	HTTP_BIND_PARAMS_ERR
)

type CustomError struct {
	ErrCode int
	ErrMsg  string
	Err     error
}

// 生成包含另一个error的error
func NewCustomError(code int, msg string) error {

	return errors.WithStack(&CustomError{code, msg, nil})
}

// 生成一个新的error
func NewCustomErrorWrap(code int, msg string, err error) error {

	return errors.WithStack(&CustomError{code, msg, err})
}

func (err *CustomError) Error() string {

	if err.Err != nil {

		return err.Err.Error()
	} else {

		return err.ErrMsg
	}
}
