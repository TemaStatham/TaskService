package handler

import (
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/middleware/auth"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getTasksUsers(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.GetTasksUsers

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	pag, err := h.takuserQuery.GetUsers(c.Request.Context(), input.TaskID, &input.Pagination, input.IsCoordinators)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"paginate": pag,
	})
}

func (h *Handler) addTasksUsers(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.AddTasksUsers

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.taskuserService.Add(c.Request.Context(), input.UserID, input.TaskID, input.IsCoordinator)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ssucc": "true",
	})
}

func (h *Handler) deleteTasksUsers(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.DeleteTasksUsers

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.taskuserService.Delete(c.Request.Context(), input.UserID, input.TaskID)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ssucc": "true",
	})
}
