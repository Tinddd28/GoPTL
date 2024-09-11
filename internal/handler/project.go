package handler

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// @Summary Create project
// @Tags projects
// @Security ApiKeyAuth
// @Description create project
// @ID create-project
// @Accept multipart/form-data
// @Produce json
// @Param title body string true "Project title"
// @Param description body string true "Project description"
// @Param token_title body string true "Token title"
// @Param amount body number true "Amount"
// @Param cost_per_token body number true "Cost per token"
// @Param image formData file true "Project image"
// @Success 200 {string} string "ok"
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /projects/create [post]
func (h *Handler) CreateProject(c *gin.Context) {
	supusr, err := getSuperUser(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !supusr {
		NewErrorResponse(c, http.StatusUnauthorized, "you are not superuser")
		return
	}

	var input models.Project
	if err := c.ShouldBind(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, "images/"+filename); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.Image = filename

	c.JSON(http.StatusOK, "ok")
}
