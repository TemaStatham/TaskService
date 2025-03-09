package main

import (
	"fmt"
	"github.com/TemaStatham/TaskService/pkg/jwt"
)

func main() {
	fmt.Println([]byte("123"))
	fmt.Println(jwt.GenerateToken(
		100002,
		"secretJWT",
	),
	)
}
