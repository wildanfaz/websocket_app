package producer

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Write(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Read
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)

		// Write
		if err := ws.WriteMessage(messageType, msg); err != nil {
			c.Logger().Error(err)
		}
	}
}
