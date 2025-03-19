package handler

import (
	"github.com/TemaStatham/TaskService/taskservice/pkg/app/response/data"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/middleware/auth"
	"github.com/TemaStatham/TaskService/taskservice/pkg/infrastructure/transport/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createResponse(c *gin.Context) {
	authUser, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.CreateResponse

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	id, err := h.responseService.Create(c.Request.Context(), input.TaskId, authUser)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"response_id": id,
	})
}

func (h *Handler) getResponses(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.GetResponses

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	// todo переписать под дто, чтобы не было необходимости вносить изменений каждый раз в этой строке
	pag, err := h.responseQuery.Show(c.Request.Context(), input.TaskId, input.Page, input.Limit)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": pag,
	})
}

func (h *Handler) rejectResponse(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.UpdateResponse

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.responseService.Update(c.Request.Context(), input.ID, "Отказано")
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func (h *Handler) confirmResponse(c *gin.Context) {
	_, err := auth.GetUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input data.UpdateResponse

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, InvalidInputBodyErr)
		return
	}

	err = h.responseService.Update(c.Request.Context(), input.ID, "Принято")
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
