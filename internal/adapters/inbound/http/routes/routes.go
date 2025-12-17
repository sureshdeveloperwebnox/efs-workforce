package routes

import (
	"efs-workforce/internal/adapters/inbound/http/handlers"
	"efs-workforce/internal/adapters/inbound/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoleRoutes sets up all role-related routes
func SetupRoleRoutes(rg *gin.RouterGroup, roleHandler *handlers.RoleHandler) {
	roles := rg.Group("/roles")
	{
		roles.POST("", roleHandler.CreateRole)
		roles.GET("", roleHandler.ListRoles)
		roles.GET("/:id", roleHandler.GetRole)
		roles.PUT("/:id", roleHandler.UpdateRole)
		roles.DELETE("/:id", roleHandler.DeleteRole)
	}
}

// SetupV1Routes sets up all v1 API routes
func SetupV1Routes(router *gin.Engine, roleHandler *handlers.RoleHandler) {
	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Apply rate limiting to API routes (if implemented)
		v1.Use(middleware.RateLimiter())

		// Role routes
		SetupRoleRoutes(v1, roleHandler)

		// TODO: Add other resource routes here
		// SetupUserRoutes(v1, userHandler)
		// SetupCrewRoutes(v1, crewHandler)
		// SetupEquipmentRoutes(v1, equipmentHandler)
	}
}
