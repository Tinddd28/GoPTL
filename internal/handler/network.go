package handler

import (
	"github.com/Tinddd28/GoPTL/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create network
// @Security ApiKeyAuth
// @Tags network
// @Description create new network
// @ID create-network
// @Accept json
// @Produce json
// @Param input body models.Network true "network info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /network [post]
func (h *Handler) CreateNetwork(c *gin.Context) {
	supus, err := getSuperUser(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !supus {
		NewErrorResponse(c, http.StatusUnauthorized, "you are not superuser")
		return
	}

	var input models.Network
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Network.CreateNetwork(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get networks
// @Tags network
// @Description get all networks
// @ID get-networks
// @Produce json
// @Success 200 {object} []models.Network
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /network [get]
func (h *Handler) GetNetworks(c *gin.Context) {
	networks, err := h.services.Network.GetNetworks()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, networks)
}

// @Summary Delete network
// @Security ApiKeyAuth
// @Tags network
// @Description delete network
// @ID delete-network
// @Param id path int true "network id"
// @Produce json
// @Success 200 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /network/{id} [delete]
func (h *Handler) DeleteNetwork(c *gin.Context) {
	supus, err := getSuperUser(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !supus {
		NewErrorResponse(c, http.StatusUnauthorized, "you are not superuser")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Network.DeleteNetwork(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
