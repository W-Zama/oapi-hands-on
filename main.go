package main

import (
	"net/http"
	"strconv"

	"example.com/oapi-hands-on/api" // ← go.mod に合わせて変更
	"github.com/labstack/echo/v4"
)

type Server struct{}

// ポインタを返すユーティリティ関数
func Ptr[T any](v T) *T {
	return &v
}

// GET /users/:id のハンドラ
func (s *Server) GetUsersId(ctx echo.Context, id int) error {
	user := api.User{
		Id:   Ptr(id),
		Name: Ptr("ユーザー " + strconv.Itoa(id)),
	}
	return ctx.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()

	// ハンドラ登録
	s := &Server{}
	api.RegisterHandlers(e, s)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}