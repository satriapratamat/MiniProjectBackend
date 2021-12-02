package presentation

import (
	"campusmanagement/features/criterias"
	"campusmanagement/features/criterias/presentation/request"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CriteriaHandler struct {
	criteriaBusiness criterias.Business
}

func NewHandlerCriteria(criteriaBusiness criterias.Business) *CriteriaHandler {
	return &CriteriaHandler{criteriaBusiness}
}

func (wlHandler *CriteriaHandler) CreateCriteria(c echo.Context) error {
	newCriteria := request.ReqCriteria{}
	id, _ := strconv.Atoi(c.Param("id"))
	newCriteria.FacultyID = id
	if err := c.Bind(&newCriteria); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := wlHandler.criteriaBusiness.CreateCriteria(newCriteria.ToCriteriaCore()); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Success",
		"data":    newCriteria,
	})
}
