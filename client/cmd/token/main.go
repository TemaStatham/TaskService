package main

import (
	"fmt"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/jwt"
)

func main() {
	fmt.Println([]byte("123"))
	fmt.Println(jwt.GenerateToken(1, "secretJWT"))
}
