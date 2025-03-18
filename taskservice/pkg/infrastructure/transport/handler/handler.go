package handler

import (
	approveservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/approve/service"
	commentquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/query"
	commentservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/comment/service"
	responsequery "github.com/TemaStatham/TaskService/taskservice/pkg/app/response/query"
	responseservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/response/service"
	taskquery "github.com/TemaStatham/TaskService/taskservice/pkg/app/task/query"
	taskservice "github.com/TemaStatham/TaskService/taskservice/pkg/app/task/service"
	hub "github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/handler/hub"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/middleware/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	InvalidInputBodyErr = "invalid input body"
)

type Handler struct {
	responseQuery   responsequery.ResponseQueryInterface
	responseService responseservice.ResponseServiceInterface
	commentQuery    commentquery.CommentQueryInterface
	commentService  commentservice.CommentServiceInterface
	taskQuery       taskquery.TaskQueryInterface
	taskService     taskservice.TaskServiceInterface
	approveService  approveservice.ApproveServiceInterface
	taskuserService taskservice.TaskUserServiceInterface
	takuserQuery    taskquery.TaskUserQueryInterface
}

func NewTaskHandler(
	responseQuery responsequery.ResponseQueryInterface,
	responseService responseservice.ResponseServiceInterface,
	commentQuery commentquery.CommentQueryInterface,
	commentService commentservice.CommentServiceInterface,
	taskQuery taskquery.TaskQueryInterface,
	taskService taskservice.TaskServiceInterface,
	approveService approveservice.ApproveServiceInterface,
	taskuserService taskservice.TaskUserServiceInterface,
	takuserQuery taskquery.TaskUserQueryInterface,
) *Handler {
	return &Handler{
		responseQuery,
		responseService,
		commentQuery,
		commentService,
		taskQuery,
		taskService,
		approveService,
		taskuserService,
		takuserQuery,
	}
}

func (h *Handler) Init(jwtSecret string) *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	httphands := router.Group("/api")
	{
		httphands.Use(auth.UserIdentity(jwtSecret))
		// Добавление, удаление, получение волонтеров или координаторов
		tasksUsers := httphands.Group("/tasks-users")
		{
			tasksUsers.GET("/byTaskID/:id", h.getTasksUsers) // получить всех возможных пользователей для задания
			tasksUsers.POST("/add/:id", h.addTasksUsers)
			tasksUsers.DELETE("/delete", h.deleteTasksUsers)
		}

		// Работа над заявками
		tasks := httphands.Group("/tasks")
		{
			tasks.GET("/all", h.getTasks)
			tasks.GET("/byID/:id", h.getTask)
			tasks.POST("/create", h.createTask)
			tasks.PUT("/update", h.updateTask)
			tasks.DELETE("/delete", h.deleteTask)
		}

		// Работа с откликами
		responses := httphands.Group("/responses")
		{
			responses.GET("/all", h.getResponses)
			responses.POST("/create", h.createResponse)
			responses.PUT("/update", h.updateResponse)
		}

		// Работа с подтверждением участия пользователя
		approves := httphands.Group("/approves")
		{
			approves.POST("/create", h.addApproves)
		}
	}

	wsHub := hub.NewHub(h.commentService, h.commentQuery)
	go wsHub.Run()

	router.GET("/ws", func(c *gin.Context) {
		roomIDParam := c.Query("roomID")

		if roomIDParam == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		roomID, err := strconv.ParseUint(roomIDParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		hub.ServeWS(c, uint(roomID), wsHub)
	})

	return router
}
