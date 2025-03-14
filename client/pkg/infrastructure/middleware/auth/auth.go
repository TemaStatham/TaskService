package auth

import (
	"errors"
	"fmt"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/jwt"
	"github.com/TemaStatham/TaskService/client/pkg/infrastructure/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func UserIdentity(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			response.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
			return
		}

		userId, err := jwt.ValidateToken(headerParts[1], jwtSecret)
		if err != nil {
			response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set(userCtx, userId.UserID)
	}
}

func GetUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	fmt.Println(id)
	idInt, ok := id.(uint)
	if !ok {
		return 0, errors.New(fmt.Sprintf("user id is of invalid type: %s", id))
	}

	return idInt, nil
}
