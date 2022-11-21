package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("P"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":3030"))
}

//**v2
// var upgrader = websocket.Upgrader{
// 	ReadBufferSize: 1024,
// 	WriteBufferSize: 1024,
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w,"home")
// }

// func read(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		conn.WriteJSON(map[string]string{"halo" : "hi"})

// 		log.Println(string(p))
// 	}
// }

// func endpoint(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool {
// 		return true
// 	}

// 	ws, err := upgrader.Upgrade(w, r, nil)

// 	defer ws.Close()

// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	read(ws)

// 	log.Println("client connected")
// }

// func routes() {
// 	http.HandleFunc("/", home)
// 	http.HandleFunc("/ws", endpoint)
// }

// func main() {
// 	routes()
// 	fmt.Println("web socket")
// 	log.Fatal(http.ListenAndServe(":3030", nil))
// }
