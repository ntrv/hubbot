package handler

import (
	"net/http"
	"os"

	chatwork "github.com/griffin-stewie/go-chatwork"
	"github.com/labstack/echo"
)

func SendChatworkPostAction(msg string, c echo.Context) error {
	cw := chatwork.NewClient(os.Getenv("API_KEY"))
	res, err := cw.PostRoomMessage(os.Getenv("ROOM_ID"), msg)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}
	return c.JSON(http.StatusOK, res)
}

func PrintPostAction(msg string, c echo.Context) error {
	return c.String(http.StatusOK, msg)
}
