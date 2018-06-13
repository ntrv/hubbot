package handler

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"

	chatwork "github.com/griffin-stewie/go-chatwork"
	"github.com/labstack/echo"
)

type PostProcessFunc func(msg string, c echo.Context) error

func SendChatworkPostProcess(msg string, c echo.Context) error {
	cw := chatwork.NewClient(os.Getenv("API_KEY"))
	res64, err := cw.PostRoomMessage(os.Getenv("ROOM_ID"), msg)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	// Cut off string before comma separator
	res64data := res64[strings.IndexByte(string(res64), ',')+1:]
	res, err := base64.URLEncoding.DecodeString(string(res64data))
	if err != nil {
		log.Println("Failed to Decode Base64: ", err.Error())
		return c.JSON(http.StatusOK, res64)
	}
	return c.JSON(http.StatusOK, res)
}

func PrintPostProcess(msg string, c echo.Context) error {
	return c.String(http.StatusOK, msg)
}
