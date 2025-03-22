package handler

import (
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/task/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/middleware/auth"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTask(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.CreateTask

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	id, err := h.taskService.Create(c.Request.Context(), &input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"task_id": id,
	})
}

func (h *Handler) updateTask(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.UpdateTask

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.taskService.Update(c.Request.Context(), &input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) deleteTask(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.DeleteTask

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.taskService.Delete(c.Request.Context(), input.ID)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func (h *Handler) getTask(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.GetTask

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	task, err := h.taskQuery.Get(c.Request.Context(), input.ID)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"task": task,
	})
}

func (h *Handler) getTasks(c *gin.Context) {
	authUser, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.GetAllTasks

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	pag, err := h.taskQuery.Show(c.Request.Context(), input, authUser)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"paginate": pag,
	})
}
