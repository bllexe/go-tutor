package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

// sha256sum
var secretKey = []byte("061eccbea6a1eb6d3b36aa7c820aa69cb0139bf95ebb5923e34b3e75f4353efb")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Add a custom claims struct for better type safety
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createToken(username string) (string, error) {
	// Use structured claims with expiration
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func verifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: In production, use secure password comparison
	if u.Username == "root" && u.Password == "123456" {
		tokenString, err := createToken(u.Username)
		if err != nil {
			http.Error(w, "Error creating token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func ProtectHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return
	}

	// Check Bearer prefix
	if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	tokenString = tokenString[7:]
	claims, err := verifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Return protected data
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Protected resource accessed successfully",
		"username": claims.Username,
	})
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/protected", ProtectHandler).Methods("GET")

	fmt.Println("Starting the server")

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		fmt.Println("Server failed to start")
	}
}
