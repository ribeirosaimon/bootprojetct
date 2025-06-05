package http

import (
	"net/http"
)

type BootContextResponse interface {
	JSON(code int, obj any)
}

// Struct principal
type response struct {
	responseContext BootContextResponse
}

// Construtor
func NewResponse(c BootContextResponse) *response {
	return &response{c}
}

// Sucesso (200)
func (r *response) Success(data any) {
	r.responseContext.JSON(http.StatusOK, data)
}

// Criado (201)
func (r *response) Created(data any) {
	r.responseContext.JSON(http.StatusCreated, data)
}

// Bad Request (400)
func (r *response) BadRequest(data any) {
	r.responseContext.JSON(http.StatusBadRequest, data)
}

// Unauthorized (401)
func (r *response) Unauthorized(data any) {
	r.responseContext.JSON(http.StatusUnauthorized, data)
}

// Forbidden (403)
func (r *response) Forbidden(data any) {
	r.responseContext.JSON(http.StatusForbidden, data)
}

// Not Found (404)
func (r *response) NotFound(data any) {
	r.responseContext.JSON(http.StatusNotFound, data)
}

// Internal Server Error (500)
func (r *response) InternalError(data any) {
	r.responseContext.JSON(http.StatusInternalServerError, data)
}

// Resposta customizada
func (r *response) Custom(status int, data any) {
	r.responseContext.JSON(status, data)
}
