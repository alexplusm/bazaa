package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	// "github.com/DATA-DOG/go-sqlmock" // todo: remove from all (.mod and .sum and others)

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	JSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestCreateGame(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(JSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	p := new(pgxpool.Pool)

	handler := CreateGame(p)

	handler(ctx)

	if assert.NoError(t, handler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	// todo
	// echo.MIMEApplicationJSON
}
