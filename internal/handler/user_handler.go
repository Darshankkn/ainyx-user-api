package handler

import (
	"ainyx-user-api/internal/models"
	"ainyx-user-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"ainyx-user-api/internal/middleware"
	"ainyx-user-api/internal/logger"
	"go.uber.org/zap"
)
import (
	"strconv"
)


type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := middleware.Validate.Struct(req); err != nil {
	return c.Status(400).JSON(fiber.Map{
		"error": err.Error(),
	})
}

	user, err := h.service.CreateUser(req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	logger.Log.Info(
	"User Created",
	zap.String("name", req.Name),
)

	return c.Status(201).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user id",
		})
	}

	user, err := h.service.GetUserByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {

	users, err := h.service.GetAllUsers()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user id",
		})
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := middleware.Validate.Struct(req); err != nil {
	return c.Status(400).JSON(fiber.Map{
		"error": err.Error(),
	})
}

	user, err := h.service.UpdateUser(id, req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
logger.Log.Info(
	"User Updated",
	zap.String("name", req.Name),
)

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid user id",
		})
	}

	err = h.service.DeleteUser(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	logger.Log.Info(
	"User Deleted",
	zap.Int("id", id),
)

	return c.SendStatus(204)
}