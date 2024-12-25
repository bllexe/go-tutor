package main


import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

func main(){
	serverAdress := "ws://localhost:8080/ws"

	conn, _, err := websocket.DefaultDialer.Dial(serverAdress,nil)
	if err != nil {
		log.Fatal("Error connecting to server: ", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	message := "Hello from client"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error sending message: ", err)
	}
	_,p,err := conn.ReadMessage()

	if err != nil{
		log.Println("Error reading message: ", err)
	}

	fmt.Printf("Received message: %s\n", string(p))

}
