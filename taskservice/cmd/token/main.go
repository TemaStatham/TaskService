package main

import (
	"fmt"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/lib/jwt"
)

func main() {
	fmt.Println([]byte("123"))
	fmt.Println(jwt.GenerateToken(1, "secretJWT"))
}
