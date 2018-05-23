package github

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestVerifyMiddleware(t *testing.T) {
	e := echo.New()
	buf := new(strings.Builder)
	e.Logger.SetOutput(buf)

	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()

	ver := VerifyMiddleware(
		VerifyConfig{
			Secret: "hogehoge",
		},
	)(echo.NotFoundHandler)

	e.POST("/", ver)

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusForbidden, rec.Code)
}
