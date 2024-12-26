package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)
//sha256
var secretKey = [] byte("061eccbea6a1eb6d3b36aa7c820aa69cb0139bf95ebb5923e34b3e75f4353efb")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func createToken(username string) (string ,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		
		"username": username,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err:=token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token,err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error){
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid{
		return fmt.Errorf("invalid token")
	}
	return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","application/json")
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v\n",u)

	if u.Username == "root" && u.Password == "123456"{
		tokenString ,err := createToken(u.Username)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No username found")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(tokenString))
		return
	}else{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Errorf("invalid credentials")
	}
}

func ProtectHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == ""{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Errorf("No token found")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err :=verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w,"Invalid token")
		return

	}

	fmt.Fprint(w,"Protected endpoint")
}

func main(){

	router := mux.NewRouter()

	router.HandleFunc("/login",LoginHandler).Methods("POST")
	router.HandleFunc("/protected",ProtectHandler).Methods("GET")

	fmt.Println("Starting the server")

	err := http.ListenAndServe("localhost:8000",router)
	if err != nil {
		fmt.Println("Server failed to start")
	}
}
