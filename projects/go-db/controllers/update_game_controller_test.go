package controllers

//import (
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//
//	"github.com/labstack/echo"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestUpdateGame(t *testing.T) {
//	var (
//		data = `{
//			"schedules": [
//				{ "schedule_id": "123" },
//				{ "schedule_id": "890" },
//				{ "schedule_id": "123" }
//			]
//		}`
//	)
//
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(data))
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//	rec := httptest.NewRecorder()
//	ctx := e.NewContext(req, rec)
//
//	p := getTestPool()
//	handler := UpdateGame(p)
//
//	if assert.NoError(t, handler(ctx)) {
//		fmt.Println()
//		assert.Equal(t, http.StatusOK, rec.Code)
//	}
//}
