package main

import (
"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection( w http.ResponseWriter, r *http.Request){

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Client connected")


	for {
		messageType, p,err := conn.ReadMessage()
		if err != nil{
			log.Println("Error reading message: ", err)
			break
		}

		err =conn.WriteMessage(messageType,p)
		if err != nil{
			log.Println("Error writing message: ", err)
			break
		}
	}

}

func main(){
	http.HandleFunc("/ws",handleConnection)

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}





