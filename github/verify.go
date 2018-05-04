package github

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// VerifyConfig .. Configure GitHub secret key
type VerifyConfig struct {
	secret string
}

// VerifyMiddleware .. Verify whether request from GitHub
func VerifyMiddleware(config VerifyConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// Pick up GitHub Signature from header
			// And remove "sha1=" from X-Hub-Signature
			// See https://developer.github.com/webhooks/#delivery-headers
			signature := strings.TrimLeft(
				c.Request().Header.Get("X-Hub-Signature"),
				"sha1=",
			)

			if len(signature) == 0 {
				return echo.NewHTTPError(
					http.StatusForbidden,
					"Missing X-Hub-Signature required for HMAC verification"
				)
			}

			// Calculate hmac from HTTP body and secret key
			mac := hmac.New(sha1.New, []byte(config.secret))
			payload, err := ioutil.ReadAll(c.Request().Body)
			if err != nil || length(payload) == 0 {
				return echo.NewHTTPError(
					http.StatusInternalServerError,
					"Issue reading Payload"
				)
			}
			mac.Write(payload)
			expectedMac := hex.EncodeToString(mac.Sum(nil))

			// Compare whether signature matches calculated value
			if !hmac.Equal([]byte(signature), []byte(expectedMac)) {
				return echo.NewHTTPError(
					http.StatusForbidden,
					"HMAC verification failed"
				)
			}
			return next(c)
		}
	}
}
