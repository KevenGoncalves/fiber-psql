package users

import (
	"database/sql"
	"github.com/KevenGoncalves/fiber-psql/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	queries *database.Queries
}

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	LastName string `json:"lastName"`
}

type UpdateUserDTO struct {
	ID       uuid.UUID
	Name     string `json:"name"`
	Email    string `json:"email"`
	LastName string `json:"lastName"`
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{
		queries: database.New(db),
	}
}

func (t *UserController) Create(c *fiber.Ctx) error {

	var body CreateUserDTO
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user database.CreateUserParams

	user.Email = &body.Email
	user.Name = &body.Name
	user.LastName = &body.LastName

	result, err := t.queries.CreateUser(c.Context(), user)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&result)
}

func (t *UserController) List(c *fiber.Ctx) error {

	result, err := t.queries.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&result)
}

func (t *UserController) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	result, err := t.queries.GetUser(c.Context(), uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&result)

}

func (t *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = t.queries.DeleteUser(c.Context(), uuid)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Deleted",
	})
}

func (t *UserController) Update(c *fiber.Ctx) error {

	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	_, err = t.queries.GetUser(c.Context(), uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	var body UpdateUserDTO
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user database.UpdateUserParams

	user.ID = uuid
	user.Email = &body.Email
	user.Name = &body.Name
	user.LastName = &body.LastName

	result, err := t.queries.UpdateUser(c.Context(), user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&result)
}
