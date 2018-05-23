package github

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestParsePayloadOnlyPOST(t *testing.T) {
	e := echo.New()
	buf := new(strings.Builder)
	e.Logger.SetOutput(buf)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	hook := NewHook()
	e.POST("/", hook.ParsePayloadHandler)
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

func TestParsePayloadNoEvent(t *testing.T) {
	e := echo.New()
	buf := new(strings.Builder)
	e.Logger.SetOutput(buf)

	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()

	hook := NewHook()
	e.POST("/", hook.ParsePayloadHandler)
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestParsePayloadNoRegister(t *testing.T) {
	e := echo.New()
	buf := new(strings.Builder)
	e.Logger.SetOutput(buf)

	req := httptest.NewRequest(echo.POST, "/", nil)
	req.Header.Set("X-GitHub-Event", "Push")
	rec := httptest.NewRecorder()

	hook := NewHook()
	e.POST("/", hook.ParsePayloadHandler)
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotImplemented, rec.Code)
}
