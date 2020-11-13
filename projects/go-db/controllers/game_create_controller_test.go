package controllers

//import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//
//	"github.com/jackc/pgx/v4/pgxpool"
//	"github.com/labstack/echo"
//	"github.com/stretchr/testify/assert"
//)
//
//func getTestPool() *pgxpool.Pool {
//	p := new(pgxpool.Pool)
//
//	// todo: docker test
//	// conn, err := dbcon.Connect()
//	// if err != nil {
//	// 	// logger
//	// 	fmt.Printf("Unable to connection to database: %v\n", err)
//	// }
//	// defer conn.Close()
//	// fmt.Println("Connected to database!")
//
//	return p
//}
//
//func TestCreateGame(t *testing.T) {
//	var (
//		data = `{
//			"name": "new game",
//			"answer_type": 1,
//			"start_date": "2020-11-02T21:33:35.298Z",
//			"end_date": "2020-11-02T21:33:35.299Z",
//			"question": "Выберите правильный ответ",
//			"options": "yep, nope"
//		}`
//	)
//
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	ctx := e.NewContext(req, rec)
//
//	p := getTestPool()
//	handler := CreateGame(p)
//
//	if assert.NoError(t, handler(ctx)) {
//		fmt.Println()
//		assert.Equal(t, http.StatusOK, rec.Code)
//	}
//}
