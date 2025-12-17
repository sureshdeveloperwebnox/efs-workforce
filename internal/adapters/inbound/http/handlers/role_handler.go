package handlers

import (
	"efs-workforce/internal/adapters/inbound/http/utils"
	"efs-workforce/internal/application"
	"efs-workforce/internal/application/dto"

	"github.com/gin-gonic/gin"
)

// RoleHandler handles role-related HTTP requests
type RoleHandler struct {
	roleService *application.RoleService
}

// NewRoleHandler creates a new role handler
func NewRoleHandler(roleService *application.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

// CreateRole handles POST /api/v1/roles
// @Summary Create a new role
// @Description Create a new role in the system
// @Tags roles
// @Accept json
// @Produce json
// @Param role body dto.CreateRoleRequest true "Role data"
// @Success 201 {object} utils.SuccessResponse{data=dto.RoleResponse}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 409 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/v1/roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req dto.CreateRoleRequest

	if !utils.BindJSON(c, &req) {
		return
	}

	resp, err := h.roleService.CreateRole(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondCreated(c, resp)
}

// GetRole handles GET /api/v1/roles/:id
// @Summary Get a role by ID
// @Description Get detailed information about a specific role
// @Tags roles
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} utils.SuccessResponse{data=dto.RoleResponse}
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/v1/roles/{id} [get]
func (h *RoleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.roleService.GetRole(id)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondOK(c, resp)
}

// ListRoles handles GET /api/v1/roles
// @Summary List all roles
// @Description Get a list of all roles in the system
// @Tags roles
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]dto.RoleResponse}
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/v1/roles [get]
func (h *RoleHandler) ListRoles(c *gin.Context) {
	roles, err := h.roleService.ListRoles()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// Return array directly for KrakenD compatibility
	c.JSON(200, roles)
}

// UpdateRole handles PUT /api/v1/roles/:id
// @Summary Update a role
// @Description Update an existing role's information
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param role body dto.UpdateRoleRequest true "Updated role data"
// @Success 200 {object} utils.SuccessResponse{data=dto.RoleResponse}
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 409 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/v1/roles/{id} [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateRoleRequest

	if !utils.BindJSON(c, &req) {
		return
	}

	resp, err := h.roleService.UpdateRole(id, &req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondOK(c, resp)
}

// DeleteRole handles DELETE /api/v1/roles/:id
// @Summary Delete a role
// @Description Delete a role from the system
// @Tags roles
// @Param id path string true "Role ID"
// @Success 204 "No Content"
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/v1/roles/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	if err := h.roleService.DeleteRole(id); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondNoContent(c)
}
