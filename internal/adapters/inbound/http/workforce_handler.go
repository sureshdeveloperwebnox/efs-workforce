package http

import (
	"efs-workforce/internal/application"
	"efs-workforce/internal/application/dto"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// WorkforceHandler implements the HTTP handler for workforce service
type WorkforceHandler struct {
	roleService *application.RoleService
	// TODO: Add other services
}

// NewWorkforceHandler creates a new HTTP handler
func NewWorkforceHandler(roleService *application.RoleService) *WorkforceHandler {
	return &WorkforceHandler{
		roleService: roleService,
	}
}

// SetupRoutes registers all HTTP routes
func (h *WorkforceHandler) SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/healthz", h.HealthCheck).Methods("GET")

	// API v1 routes
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	// Role routes
	apiV1.HandleFunc("/roles", h.CreateRole).Methods("POST")
	apiV1.HandleFunc("/roles/{id}", h.GetRole).Methods("GET")
	apiV1.HandleFunc("/roles", h.ListRoles).Methods("GET")
	apiV1.HandleFunc("/roles/{id}", h.UpdateRole).Methods("PUT")
	apiV1.HandleFunc("/roles/{id}", h.DeleteRole).Methods("DELETE")

	// TODO: Add routes for other entities

	return router
}

// HealthCheck handles GET /healthz
func (h *WorkforceHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "healthy",
		"service": "workforce-service",
	})
}

// CreateRole handles POST /api/v1/roles
func (h *WorkforceHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.roleService.CreateRole(&req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GetRole handles GET /api/v1/roles/{id}
func (h *WorkforceHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := h.roleService.GetRole(id)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// ListRoles handles GET /api/v1/roles
func (h *WorkforceHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.roleService.ListRoles()
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// UpdateRole handles PUT /api/v1/roles/{id}
func (h *WorkforceHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req dto.UpdateRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.roleService.UpdateRole(id, &req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DeleteRole handles DELETE /api/v1/roles/{id}
func (h *WorkforceHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.roleService.DeleteRole(id); err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// handleError handles errors and returns appropriate HTTP responses
func (h *WorkforceHandler) handleError(w http.ResponseWriter, err error) {
	if strings.Contains(err.Error(), "already exists") {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	// In a real implementation, you would map domain errors to HTTP status codes
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
