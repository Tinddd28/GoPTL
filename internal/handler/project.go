package handler

import (
	"fmt"
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create a new project
// @Tags projects
// @Security ApiKeyAuth
// @Description Create a new project with title, description, token title, amount, cost per token, and an image.
// @Accept  multipart/form-data
// @Produce json
// @Param title formData string true "Project title"
// @Param description formData string true "Project description"
// @Param token_title formData string true "Token title"
// @Param amount formData number true "Amount of tokens"
// @Param cost_per_token formData number true "Cost per token"
// @Param image formData file true "Project image file"
// @Success 200 {object} models.Project
// @Failure 400 {object} ErrorResponse "Invalid request data"
// @Failure 500 {object} ErrorResponse "Internal server error"
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

	//var project models.Project
	//
	//project.Title = c.PostForm("title")
	//project.Description = c.PostForm("description")
	//project.TokenTitle = c.PostForm("token_title")
	//amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	//if err != nil {
	//	NewErrorResponse(c, http.StatusBadRequest, "invalid amount")
	//	return
	//}
	//
	//project.Amount = amount
	//
	//costPerToken, err := strconv.ParseFloat(c.PostForm("cost_per_token"), 64)
	//if err != nil {
	//	NewErrorResponse(c, http.StatusBadRequest, "invalid cost per token")
	//	return
	//}
	//
	//project.CostPerToken = costPerToken
	var form models.ProjectForm

	// Привязываем данные формы
	if err := c.ShouldBind(&form); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "image upload error")
		return
	}

	// Сохраняем файл
	imagePath := fmt.Sprintf("images/%s", file.Filename)
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "failed to save image")
		return
	}

	project := models.Project{
		Title:        form.Title,
		Description:  form.Description,
		TokenTitle:   form.TokenTitle,
		Amount:       form.Amount,
		CostPerToken: form.CostPerToken,
		Image:        imagePath,
	}

	id, err := h.services.Project.CreateProject(project)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get projects
// @Tags projects
// @Description get projects
// @ID get-projects
// @Param offset query int true "offest"
// @Produce json
// @Success 200 {object} []models.Project
// @Failure 500 {object} ErrorResponse
// @Router /projects/all [get]
func (h *Handler) GetProjects(c *gin.Context) {

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid query param")
		return
	}
	projects, err := h.services.Project.GetProjects(offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, projects)
}

// @Summary Get project by id
// @Tags projects
// @Description get project by id
// @ID get-project-by-id
// @Param id path int true "Project id"
// @Produce json
// @Success 200 {object} models.Project
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /projects/{id} [get]
func (h *Handler) GetProjectById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	project, err := h.services.Project.GetProjectById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, project)
}

// @Summary Update project
// @Tags projects
// @Security ApiKeyAuth
// @Description update project
// @ID update-project
// @Param id path int true "Project id"
// @Accept json
// @Produce json
// @Param title body string true "Project title"
// @Param description body string true "Project description"
// @Param token_title body string true "Token title"
// @Param amount body number true "Amount"
// @Param cost_per_token body number true "Cost per token"
// @Success 200 {object} string
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /projects/{id} [put]
func (h *Handler) UpdateProject(c *gin.Context) {
	supusr, err := getSuperUser(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !supusr {
		NewErrorResponse(c, http.StatusUnauthorized, "you are not superuser")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.Project
	if err := c.ShouldBind(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Project.UpdateProject(id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// @Summary Delete project
// @Tags projects
// @Security ApiKeyAuth
// @Description delete project
// @ID delete-project
// @Param id path int true "Project id"
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /projects/{id} [delete]
func (h *Handler) DeleteProject(c *gin.Context) {
	supusr, err := getSuperUser(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !supusr {
		NewErrorResponse(c, http.StatusUnauthorized, "you are not superuser")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Project.DeleteProject(id); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
