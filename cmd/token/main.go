package main

import (
	"fmt"
	"github.com/TemaStatham/TaskService/internal/model"
	"github.com/TemaStatham/TaskService/pkg/jwt"
	"time"
)

func main() {
	fmt.Println([]byte("123"))
	fmt.Println(jwt.NewToken(
		model.User{
			ID:       1,
			Email:    "admin@tema.com",
			PassHash: []byte("123"),
			Role:     1,
		},
		time.Minute*300,
	),
	)
}
