package main

import (
	"fmt"
	"time"
)

// 独自定義のエラー構造体
type MyError struct {
	Code    int
	Message string
	When    time.Time
}

// errorインターフェースのメソッドを定義
func (e *MyError) Error() string {
	return fmt.Sprintf("%v [エラーコード: %d] %s",
		e.When, e.Code, e.Message)
}

// エラーを発生させる関数
func run() error {
	return &MyError{
		Code:    8888,
		Message: "エラーが発生しました",
		When:    time.Now(),
	}
}

func main() {
	// エラーを発生させる
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
