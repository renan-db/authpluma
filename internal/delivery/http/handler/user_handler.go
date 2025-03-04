package http

import (
	"net/http"
	database "project/internal/infrastructure/database/sqlc"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	querier database.Querier
}

func NewUserHandler(querier database.Querier) *UserHandler {
	return &UserHandler{querier: querier}
}

// CreateUser godoc
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	params := database.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	}

	createdUser, err := h.querier.CreateUser(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

// GetUser godoc
func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	user, err := h.querier.GetUser(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

// ListUsers godoc
func (h *UserHandler) ListUsers(c echo.Context) error {
	users, err := h.querier.ListUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser godoc
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	params := database.UpdateUserParams{
		ID:    int32(id),
		Name:  user.Name,
		Email: user.Email,
	}

	updatedUser, err := h.querier.UpdateUser(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	err = h.querier.DeleteUser(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
} 